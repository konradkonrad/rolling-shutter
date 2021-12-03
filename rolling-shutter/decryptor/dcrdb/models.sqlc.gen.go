// Code generated by sqlc. DO NOT EDIT.

package dcrdb

import ()

type DecryptorAggregatedSignature struct {
	EpochID         []byte
	SignedHash      []byte
	SignersBitfield []byte
	Signature       []byte
}

type DecryptorChainCollator struct {
	ActivationBlockNumber int64
	Collator              string
}

type DecryptorChainKeyperSet struct {
	N         int32
	Addresses []string
}

type DecryptorCipherBatch struct {
	EpochID      []byte
	Transactions [][]byte
}

type DecryptorDecryptionKey struct {
	EpochID []byte
	Key     []byte
}

type DecryptorDecryptionSignature struct {
	EpochID         []byte
	SignedHash      []byte
	SignersBitfield []byte
	Signature       []byte
}

type DecryptorDecryptorIdentity struct {
	Address           string
	BlsPublicKey      []byte
	BlsSignature      []byte
	SignatureVerified bool
}

type DecryptorDecryptorSetMember struct {
	ActivationBlockNumber int64
	Index                 int32
	Address               string
}

type DecryptorEonPublicKey struct {
	ActivationBlockNumber int64
	EonPublicKey          []byte
}

type DecryptorEventSyncProgress struct {
	ID              bool
	NextBlockNumber int32
	NextLogIndex    int32
}

type DecryptorKeyperSet struct {
	ActivationBlockNumber int64
	Keypers               []string
	Threshold             int32
}

type DecryptorMetaInf struct {
	Key   string
	Value string
}
