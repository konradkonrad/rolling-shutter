package decryptor

import (
	"context"
	"fmt"
	"log"
	"math"

	"github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"

	"github.com/shutter-network/shutter/shlib/shcrypto"
	"github.com/shutter-network/shutter/shlib/shcrypto/shbls"
	"github.com/shutter-network/shutter/shuttermint/decryptor/dcrdb"
	"github.com/shutter-network/shutter/shuttermint/medley"
	"github.com/shutter-network/shutter/shuttermint/shmsg"
)

func handleDecryptionKeyInput(
	ctx context.Context,
	config Config,
	db *dcrdb.Queries,
	key *decryptionKey,
) ([]shmsg.P2PMessage, error) {
	eonPublicKeyBytes, err := db.GetEonPublicKey(ctx, medley.Uint64EpochIDToBytes(key.epochID))
	if err == pgx.ErrNoRows {
		return nil, errors.Errorf(
			"received decryption key for epoch %d for which we don't have an eon public key",
			key.epochID,
		)
	}
	if err != nil {
		return nil, err
	}
	eonPublicKey := new(shcrypto.EonPublicKey)
	err = eonPublicKey.Unmarshal(eonPublicKeyBytes)
	if err != nil {
		return nil, err
	}
	ok, err := checkEpochSecretKey(key.key, eonPublicKey, key.epochID)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.Errorf("received decryption key does not match eon public key for epoch %d", key.epochID)
	}

	keyBytes, _ := key.key.GobEncode()
	tag, err := db.InsertDecryptionKey(ctx, dcrdb.InsertDecryptionKeyParams{
		EpochID: medley.Uint64EpochIDToBytes(key.epochID),
		Key:     keyBytes,
	})
	if err != nil {
		return nil, err
	}
	if tag.RowsAffected() == 0 {
		log.Printf("attempted to store multiple keys for same epoch %d", key.epochID)
		return nil, nil
	}
	return handleEpoch(ctx, config, db, key.epochID)
}

func handleCipherBatchInput(
	ctx context.Context,
	config Config,
	db *dcrdb.Queries,
	cipherBatch *cipherBatch,
) ([]shmsg.P2PMessage, error) {
	tag, err := db.InsertCipherBatch(ctx, dcrdb.InsertCipherBatchParams{
		EpochID:      medley.Uint64EpochIDToBytes(cipherBatch.EpochID),
		Transactions: cipherBatch.Transactions,
	})
	if err != nil {
		return nil, err
	}
	if tag.RowsAffected() == 0 {
		log.Printf("attempted to store multiple cipherbatches for same epoch %d", cipherBatch.EpochID)
		return nil, nil
	}
	return handleEpoch(ctx, config, db, cipherBatch.EpochID)
}

func handleSignatureInput(
	ctx context.Context,
	config Config,
	db *dcrdb.Queries,
	signature *decryptionSignature,
) ([]shmsg.P2PMessage, error) {
	signers := getSignerIndexes(signature.SignerBitfield)
	if len(signers) > 1 {
		// Ignore aggregated signatures
		return nil, nil
	}
	tag, err := db.InsertDecryptionSignature(ctx, dcrdb.InsertDecryptionSignatureParams{
		EpochID:         medley.Uint64EpochIDToBytes(signature.epochID),
		SignedHash:      signature.signedHash.Bytes(),
		SignersBitfield: signature.SignerBitfield,
		Signature:       signature.signature.Marshal(),
	})
	if err != nil {
		return nil, err
	}
	if tag.RowsAffected() == 0 {
		log.Printf("attempted to store multiple decryption signatures with same epoch and signers")
		return nil, nil
	}

	// check if we have enough signatures
	dbSignatures, err := db.GetDecryptionSignatures(ctx, medley.Uint64EpochIDToBytes(signature.epochID))
	if err != nil {
		return nil, err
	}
	if uint(len(dbSignatures)) < config.RequiredSignatures {
		return nil, nil
	}

	signaturesToAggregate := make([]*shbls.Signature, 0, len(dbSignatures))
	publicKeysToAggragate := make([]*shbls.PublicKey, 0, len(dbSignatures))
	bitfield := make([]byte, len(signature.SignerBitfield))
	for _, dbSignature := range dbSignatures {
		if common.BytesToHash(dbSignature.SignedHash) != signature.signedHash {
			continue
		}

		unmarshalledSignature := new(shbls.Signature)
		if err := unmarshalledSignature.Unmarshal(dbSignature.Signature); err != nil {
			return nil, err
		}
		signaturesToAggregate = append(signaturesToAggregate, unmarshalledSignature)

		indexes := getSignerIndexes(dbSignature.SignersBitfield)
		if len(indexes) > 1 {
			panic("got signature with multiple signers")
		}
		if len(indexes) == 0 {
			panic("could not retrieve signer index from bitfield")
		}
		pkBytes, err := db.GetDecryptorKey(ctx, dcrdb.GetDecryptorKeyParams{Index: indexes[0], StartEpochID: dbSignature.EpochID})
		if err != nil {
			return nil, err
		}
		pk := new(shbls.PublicKey)
		if err := pk.Unmarshal(pkBytes); err != nil {
			return nil, err
		}
		publicKeysToAggragate = append(publicKeysToAggragate, pk)
		addBitfields(bitfield, dbSignature.SignersBitfield)
	}

	if uint(len(signaturesToAggregate)) < config.RequiredSignatures {
		return nil, nil
	}

	aggregatedSignature := shbls.AggregateSignatures(signaturesToAggregate)
	aggregatedKey := shbls.AggregatePublicKeys(publicKeysToAggragate)
	if !shbls.Verify(aggregatedSignature, aggregatedKey, signature.signedHash.Bytes()) {
		panic(fmt.Sprintf("could not verify aggregated signature for epochID %d", signature.epochID))
	}

	msgs := []shmsg.P2PMessage{
		&shmsg.AggregatedDecryptionSignature{
			InstanceID:          config.InstanceID,
			EpochID:             signature.epochID,
			SignedHash:          signature.signedHash.Bytes(),
			AggregatedSignature: aggregatedSignature.Marshal(),
			SignerBitfield:      bitfield,
		},
	}

	return msgs, nil
}

// handleEpoch produces, store, and output a signature if we have both the cipher batch and key for given epoch.
func handleEpoch(
	ctx context.Context,
	config Config,
	db *dcrdb.Queries,
	epochID uint64,
) ([]shmsg.P2PMessage, error) {
	epochIDBytes := medley.Uint64EpochIDToBytes(epochID)
	cipherBatch, err := db.GetCipherBatch(ctx, epochIDBytes)
	if err == pgx.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	decryptionKeyDB, err := db.GetDecryptionKey(ctx, epochIDBytes)
	if err == pgx.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	log.Printf("decrypting batch for epoch %d", epochID)

	decryptionKey := new(shcrypto.EpochSecretKey)
	err = decryptionKey.GobDecode(decryptionKeyDB.Key)
	if err != nil {
		return nil, errors.Wrapf(err, "invalid decryption key for epoch %d in db", epochID)
	}

	decryptedBatch := decryptCipherBatch(cipherBatch.Transactions, decryptionKey)
	signingData := DecryptionSigningData{
		InstanceID:     config.InstanceID,
		EpochID:        epochID,
		CipherBatch:    cipherBatch.Transactions,
		DecryptedBatch: decryptedBatch,
	}
	signedHash := signingData.Hash().Bytes()
	signatureBytes := signingData.Sign(config.SigningKey).Marshal()
	signerIndexes := make([]byte, 0) // TODO: find this value

	insertParams := dcrdb.InsertDecryptionSignatureParams{
		EpochID:         epochIDBytes,
		SignedHash:      signedHash,
		SignersBitfield: signerIndexes,
		Signature:       signatureBytes,
	}
	tag, err := db.InsertDecryptionSignature(ctx, insertParams)
	if err != nil {
		return nil, err
	}
	if tag.RowsAffected() == 0 {
		log.Printf("attempted to store multiple signatures with same (epoch id, signer index): (%d, %d)",
			epochID, signerIndexes)
		return nil, nil
	}

	msgs := []shmsg.P2PMessage{}
	// TODO: handle signer bitfield
	msgs = append(msgs, &shmsg.AggregatedDecryptionSignature{
		InstanceID:          config.InstanceID,
		EpochID:             epochID,
		SignedHash:          signedHash,
		AggregatedSignature: signatureBytes,
		SignerBitfield:      signerIndexes,
	})
	return msgs, nil
}

func getSignerIndexes(bitField []byte) []int32 {
	numBytes := int32(0)
	var indexes []int32
	for _, b := range bitField {
		for i := 7; i >= 0; i-- {
			threshold := uint8(math.Pow(2, float64(i)))
			if b >= threshold {
				b -= threshold
				indexes = append(indexes, numBytes*8+int32(i)+1)
			}
		}
	}
	return indexes
}

func addBitfields(bf1 []byte, bf2 []byte) []byte {
	if len(bf1) != len(bf2) {
		panic("require two bitfields of the same length")
	}
	for i, b := range bf2 {
		bf1[i] += b
	}
	return bf1
}
