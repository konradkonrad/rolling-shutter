-- name: GetCipherBatch :one
SELECT * FROM decryptor.cipher_batch
WHERE epoch_id = $1;

-- name: InsertCipherBatch :execresult
INSERT INTO decryptor.cipher_batch (
    epoch_id, transactions
) VALUES (
    $1, $2
)
ON CONFLICT DO NOTHING;

-- name: GetDecryptionKey :one
SELECT * FROM decryptor.decryption_key
WHERE epoch_id = $1;

-- name: InsertDecryptionKey :execresult
INSERT INTO decryptor.decryption_key (
    epoch_id, key
) VALUES (
    $1, $2
)
ON CONFLICT DO NOTHING;

-- name: GetDecryptionSignatures :many
SELECT * FROM decryptor.decryption_signature
WHERE epoch_id = $1 AND signed_hash = $2;

-- name: GetDecryptionSignature :one
SELECT * FROM decryptor.decryption_signature
WHERE epoch_id = $1 AND signers_bitfield = $2;

-- name: InsertDecryptionSignature :execresult
INSERT INTO decryptor.decryption_signature (
    epoch_id, signed_hash, signers_bitfield, signature
) VALUES (
    $1, $2, $3, $4
)
ON CONFLICT DO NOTHING;

-- name: GetAggregatedSignature :one
SELECT * FROM decryptor.aggregated_signature
WHERE signed_hash = $1;

-- name: ExistsAggregatedSignature :one
SELECT EXISTS(SELECT 1 FROM decryptor.aggregated_signature WHERE signed_hash = $1);

-- name: InsertAggregatedSignature :execresult
INSERT INTO decryptor.aggregated_signature (
    epoch_id, signed_hash, signers_bitfield, signature
) VALUES (
    $1, $2, $3, $4
)
ON CONFLICT DO NOTHING;

-- name: InsertDecryptorIdentity :exec
INSERT INTO decryptor.decryptor_identity (
    address, bls_public_key, bls_signature, signature_valid
) VALUES (
    $1, $2, $3, $4
);

-- name: GetDecryptorIdentity :one
SELECT * FROM decryptor.decryptor_identity
WHERE address = $1;

-- name: InsertDecryptorSetMember :exec
INSERT INTO decryptor.decryptor_set_member (
    activation_block_number, index, address
) VALUES (
    $1, $2, $3
);

-- name: GetDecryptorSetMember :one
SELECT
    m1.activation_block_number,
    m1.index,
    m1.address,
    identity.bls_public_key,
    identity.bls_signature,
    coalesce(identity.signature_valid, false)
FROM (
    SELECT
        m2.activation_block_number,
        m2.index,
        m2.address
    FROM decryptor.decryptor_set_member AS m2
    WHERE activation_block_number = (
        SELECT
            m3.activation_block_number
        FROM decryptor.decryptor_set_member AS m3
        WHERE m3.activation_block_number <= $1
        ORDER BY m3.activation_block_number DESC
        LIMIT 1
    ) AND m2.index = $2
) AS m1
LEFT OUTER JOIN decryptor.decryptor_identity AS identity
ON m1.address = identity.address
ORDER BY index;

-- name: GetDecryptorSet :many
SELECT
    member.activation_block_number,
    member.index,
    member.address,
    identity.bls_public_key,
    identity.bls_signature,
    coalesce(identity.signature_valid, false)
FROM (
    SELECT
        activation_block_number,
        index,
        address
    FROM decryptor.decryptor_set_member
    WHERE activation_block_number = (
        SELECT
            m.activation_block_number
        FROM decryptor.decryptor_set_member AS m
        WHERE m.activation_block_number <= $1
        ORDER BY m.activation_block_number DESC
        LIMIT 1
    )
) AS member
LEFT OUTER JOIN decryptor.decryptor_identity AS identity
ON member.address = identity.address
ORDER BY index;

-- name: InsertEonPublicKey :exec
INSERT INTO decryptor.eon_public_key (
    activation_block_number,
    eon_public_key
) VALUES (
    $1, $2
);

-- name: GetEonPublicKey :one
SELECT eon_public_key
FROM decryptor.eon_public_key
WHERE activation_block_number <= $1
ORDER BY activation_block_number DESC LIMIT 1;

-- name: InsertKeyperSet :exec
INSERT INTO decryptor.keyper_set (
    activation_block_number,
    keypers,
    threshold
) VALUES (
    $1, $2, $3
);

-- name: GetKeyperSet :one
SELECT (
    activation_block_number,
    keypers,
    threshold
) FROM decryptor.keyper_set
WHERE activation_block_number <= $1
ORDER BY activation_block_number DESC LIMIT 1;

-- name: InsertMeta :exec
INSERT INTO decryptor.meta_inf (key, value) VALUES ($1, $2);

-- name: GetMeta :one
SELECT * FROM decryptor.meta_inf WHERE key = $1;

-- name: UpdateEventSyncProgress :exec
INSERT INTO decryptor.event_sync_progress (next_block_number, next_log_index)
VALUES ($1, $2)
ON CONFLICT (id) DO UPDATE
    SET next_block_number = $1,
        next_log_index = $2;

-- name: GetEventSyncProgress :one
SELECT * FROM decryptor.event_sync_progress LIMIT 1;

-- name: InsertChainKeyperSet :exec
INSERT INTO decryptor.chain_keyper_set (n, addresses) VALUES ($1, $2);

-- name: GetChainKeyperSet :one
SELECT * FROM decryptor.chain_keyper_set LIMIT 1;

-- name: InsertChainCollator :exec
INSERT INTO decryptor.chain_collator (activation_block_number, collator)
VALUES ($1, $2);

-- name: GetChainCollator :one
SELECT * FROM decryptor.chain_collator
WHERE activation_block_number <= $1
ORDER BY activation_block_number DESC LIMIT 1;
