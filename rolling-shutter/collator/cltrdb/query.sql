-- name: InsertMeta :exec
INSERT INTO collator.meta_inf (key, value) VALUES ($1, $2);

-- name: GetMeta :one
SELECT * FROM collator.meta_inf WHERE key = $1;

-- name: InsertTrigger :exec
INSERT INTO collator.decryption_trigger (epoch_id, batch_hash) VALUES ($1, $2);

-- name: GetTrigger :one
SELECT * FROM collator.decryption_trigger WHERE epoch_id = $1;

-- name: GetLastTrigger :one
SELECT * FROM collator.decryption_trigger ORDER BY epoch_id DESC LIMIT 1;

-- name: InsertBatch :exec
INSERT INTO collator.cipher_batch (epoch_id, transactions) VALUES ($1, $2);

-- name: GetBatch :one
SELECT * FROM collator.cipher_batch WHERE epoch_id = $1;