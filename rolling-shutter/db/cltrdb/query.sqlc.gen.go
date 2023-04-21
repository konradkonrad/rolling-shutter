// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: query.sql

package cltrdb

import (
	"context"

	"github.com/jackc/pgconn"
)

const confirmEonPublicKey = `-- name: ConfirmEonPublicKey :exec
UPDATE eon_public_key_candidate
SET confirmed=TRUE
WHERE hash=$1
`

func (q *Queries) ConfirmEonPublicKey(ctx context.Context, hash []byte) error {
	_, err := q.db.Exec(ctx, confirmEonPublicKey, hash)
	return err
}

const countEonPublicKeyVotes = `-- name: CountEonPublicKeyVotes :one
SELECT COUNT(*) from eon_public_key_vote WHERE hash=$1
`

func (q *Queries) CountEonPublicKeyVotes(ctx context.Context, hash []byte) (int64, error) {
	row := q.db.QueryRow(ctx, countEonPublicKeyVotes, hash)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const existsDecryptionKey = `-- name: ExistsDecryptionKey :one
SELECT EXISTS (
    SELECT 1
    FROM decryption_key
    WHERE epoch_id = $1
)
`

func (q *Queries) ExistsDecryptionKey(ctx context.Context, epochID []byte) (bool, error) {
	row := q.db.QueryRow(ctx, existsDecryptionKey, epochID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const findEonPublicKeyForBlock = `-- name: FindEonPublicKeyForBlock :one
SELECT hash, eon_public_key, activation_block_number, keyper_config_index, eon, confirmed FROM eon_public_key_candidate
WHERE confirmed AND activation_block_number <= $1
ORDER BY activation_block_number DESC, keyper_config_index DESC
LIMIT 1
`

func (q *Queries) FindEonPublicKeyForBlock(ctx context.Context, blocknumber int64) (EonPublicKeyCandidate, error) {
	row := q.db.QueryRow(ctx, findEonPublicKeyForBlock, blocknumber)
	var i EonPublicKeyCandidate
	err := row.Scan(
		&i.Hash,
		&i.EonPublicKey,
		&i.ActivationBlockNumber,
		&i.KeyperConfigIndex,
		&i.Eon,
		&i.Confirmed,
	)
	return i, err
}

const findEonPublicKeyVotes = `-- name: FindEonPublicKeyVotes :many
SELECT hash, sender, signature, eon, keyper_config_index FROM eon_public_key_vote WHERE hash=$1 ORDER BY sender
`

func (q *Queries) FindEonPublicKeyVotes(ctx context.Context, hash []byte) ([]EonPublicKeyVote, error) {
	rows, err := q.db.Query(ctx, findEonPublicKeyVotes, hash)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []EonPublicKeyVote
	for rows.Next() {
		var i EonPublicKeyVote
		if err := rows.Scan(
			&i.Hash,
			&i.Sender,
			&i.Signature,
			&i.Eon,
			&i.KeyperConfigIndex,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCommittedTransactionsByEpoch = `-- name: GetCommittedTransactionsByEpoch :many
SELECT tx_hash, id, epoch_id, tx_bytes, status FROM transaction WHERE status = 'committed' AND epoch_id = $1 ORDER BY id ASC
`

func (q *Queries) GetCommittedTransactionsByEpoch(ctx context.Context, epochID []byte) ([]Transaction, error) {
	rows, err := q.db.Query(ctx, getCommittedTransactionsByEpoch, epochID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transaction
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.TxHash,
			&i.ID,
			&i.EpochID,
			&i.TxBytes,
			&i.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getDecryptionKey = `-- name: GetDecryptionKey :one
SELECT epoch_id, decryption_key FROM decryption_key
WHERE epoch_id = $1
`

func (q *Queries) GetDecryptionKey(ctx context.Context, epochID []byte) (DecryptionKey, error) {
	row := q.db.QueryRow(ctx, getDecryptionKey, epochID)
	var i DecryptionKey
	err := row.Scan(&i.EpochID, &i.DecryptionKey)
	return i, err
}

const getEonPublicKey = `-- name: GetEonPublicKey :one
SELECT hash, eon_public_key, activation_block_number, keyper_config_index, eon, confirmed FROM eon_public_key_candidate
WHERE confirmed AND eon = $1
LIMIT 1
`

func (q *Queries) GetEonPublicKey(ctx context.Context, eon int64) (EonPublicKeyCandidate, error) {
	row := q.db.QueryRow(ctx, getEonPublicKey, eon)
	var i EonPublicKeyCandidate
	err := row.Scan(
		&i.Hash,
		&i.EonPublicKey,
		&i.ActivationBlockNumber,
		&i.KeyperConfigIndex,
		&i.Eon,
		&i.Confirmed,
	)
	return i, err
}

const getLastBatchEpochID = `-- name: GetLastBatchEpochID :one
SELECT epoch_id FROM decryption_trigger ORDER BY epoch_id DESC LIMIT 1
`

func (q *Queries) GetLastBatchEpochID(ctx context.Context) ([]byte, error) {
	row := q.db.QueryRow(ctx, getLastBatchEpochID)
	var epoch_id []byte
	err := row.Scan(&epoch_id)
	return epoch_id, err
}

const getNextBatch = `-- name: GetNextBatch :one
SELECT enforce_one_row, epoch_id, l1_block_number FROM next_batch LIMIT 1
`

func (q *Queries) GetNextBatch(ctx context.Context) (NextBatch, error) {
	row := q.db.QueryRow(ctx, getNextBatch)
	var i NextBatch
	err := row.Scan(&i.EnforceOneRow, &i.EpochID, &i.L1BlockNumber)
	return i, err
}

const getNonRejectedTransactionsByEpoch = `-- name: GetNonRejectedTransactionsByEpoch :many
SELECT tx_hash, id, epoch_id, tx_bytes, status FROM transaction WHERE status<>'rejected' AND epoch_id = $1 ORDER BY id ASC
`

func (q *Queries) GetNonRejectedTransactionsByEpoch(ctx context.Context, epochID []byte) ([]Transaction, error) {
	rows, err := q.db.Query(ctx, getNonRejectedTransactionsByEpoch, epochID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transaction
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.TxHash,
			&i.ID,
			&i.EpochID,
			&i.TxBytes,
			&i.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTransactionsByEpoch = `-- name: GetTransactionsByEpoch :many
SELECT tx_hash, id, epoch_id, tx_bytes, status FROM transaction WHERE epoch_id = $1 ORDER BY id ASC
`

func (q *Queries) GetTransactionsByEpoch(ctx context.Context, epochID []byte) ([]Transaction, error) {
	rows, err := q.db.Query(ctx, getTransactionsByEpoch, epochID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transaction
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.TxHash,
			&i.ID,
			&i.EpochID,
			&i.TxBytes,
			&i.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTrigger = `-- name: GetTrigger :one
SELECT epoch_id, id, batch_hash, l1_block_number, sent FROM decryption_trigger WHERE epoch_id = $1
`

func (q *Queries) GetTrigger(ctx context.Context, epochID []byte) (DecryptionTrigger, error) {
	row := q.db.QueryRow(ctx, getTrigger, epochID)
	var i DecryptionTrigger
	err := row.Scan(
		&i.EpochID,
		&i.ID,
		&i.BatchHash,
		&i.L1BlockNumber,
		&i.Sent,
	)
	return i, err
}

const getUnsentTriggers = `-- name: GetUnsentTriggers :many
SELECT epoch_id, id, batch_hash, l1_block_number, sent FROM decryption_trigger
WHERE sent IS NULL
ORDER BY id ASC
`

func (q *Queries) GetUnsentTriggers(ctx context.Context) ([]DecryptionTrigger, error) {
	rows, err := q.db.Query(ctx, getUnsentTriggers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []DecryptionTrigger
	for rows.Next() {
		var i DecryptionTrigger
		if err := rows.Scan(
			&i.EpochID,
			&i.ID,
			&i.BatchHash,
			&i.L1BlockNumber,
			&i.Sent,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUnsubmittedBatchTx = `-- name: GetUnsubmittedBatchTx :one
SELECT epoch_id, marshaled, submitted FROM batchtx WHERE submitted=false
`

func (q *Queries) GetUnsubmittedBatchTx(ctx context.Context) (Batchtx, error) {
	row := q.db.QueryRow(ctx, getUnsubmittedBatchTx)
	var i Batchtx
	err := row.Scan(&i.EpochID, &i.Marshaled, &i.Submitted)
	return i, err
}

const insertBatchTx = `-- name: InsertBatchTx :exec
INSERT INTO batchtx (epoch_id, marshaled) VALUES ($1, $2)
`

type InsertBatchTxParams struct {
	EpochID   []byte
	Marshaled []byte
}

func (q *Queries) InsertBatchTx(ctx context.Context, arg InsertBatchTxParams) error {
	_, err := q.db.Exec(ctx, insertBatchTx, arg.EpochID, arg.Marshaled)
	return err
}

const insertDecryptionKey = `-- name: InsertDecryptionKey :execresult
INSERT INTO decryption_key (epoch_id, decryption_key)
VALUES ($1, $2)
ON CONFLICT DO NOTHING
`

type InsertDecryptionKeyParams struct {
	EpochID       []byte
	DecryptionKey []byte
}

func (q *Queries) InsertDecryptionKey(ctx context.Context, arg InsertDecryptionKeyParams) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, insertDecryptionKey, arg.EpochID, arg.DecryptionKey)
}

const insertEonPublicKeyCandidate = `-- name: InsertEonPublicKeyCandidate :exec
INSERT INTO eon_public_key_candidate
       (hash, eon_public_key, activation_block_number, keyper_config_index, eon)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT DO NOTHING
`

type InsertEonPublicKeyCandidateParams struct {
	Hash                  []byte
	EonPublicKey          []byte
	ActivationBlockNumber int64
	KeyperConfigIndex     int64
	Eon                   int64
}

func (q *Queries) InsertEonPublicKeyCandidate(ctx context.Context, arg InsertEonPublicKeyCandidateParams) error {
	_, err := q.db.Exec(ctx, insertEonPublicKeyCandidate,
		arg.Hash,
		arg.EonPublicKey,
		arg.ActivationBlockNumber,
		arg.KeyperConfigIndex,
		arg.Eon,
	)
	return err
}

const insertEonPublicKeyVote = `-- name: InsertEonPublicKeyVote :exec
INSERT INTO eon_public_key_vote
       (hash, sender, signature, eon, keyper_config_index)
VALUES ($1, $2, $3, $4, $5)
`

type InsertEonPublicKeyVoteParams struct {
	Hash              []byte
	Sender            string
	Signature         []byte
	Eon               int64
	KeyperConfigIndex int64
}

func (q *Queries) InsertEonPublicKeyVote(ctx context.Context, arg InsertEonPublicKeyVoteParams) error {
	_, err := q.db.Exec(ctx, insertEonPublicKeyVote,
		arg.Hash,
		arg.Sender,
		arg.Signature,
		arg.Eon,
		arg.KeyperConfigIndex,
	)
	return err
}

const insertTrigger = `-- name: InsertTrigger :exec
INSERT INTO decryption_trigger (epoch_id, batch_hash, l1_block_number) VALUES ($1, $2, $3)
`

type InsertTriggerParams struct {
	EpochID       []byte
	BatchHash     []byte
	L1BlockNumber int64
}

func (q *Queries) InsertTrigger(ctx context.Context, arg InsertTriggerParams) error {
	_, err := q.db.Exec(ctx, insertTrigger, arg.EpochID, arg.BatchHash, arg.L1BlockNumber)
	return err
}

const insertTx = `-- name: InsertTx :exec
INSERT INTO transaction (tx_hash, epoch_id, tx_bytes, status) VALUES ($1, $2, $3, $4)
`

type InsertTxParams struct {
	TxHash  []byte
	EpochID []byte
	TxBytes []byte
	Status  Txstatus
}

func (q *Queries) InsertTx(ctx context.Context, arg InsertTxParams) error {
	_, err := q.db.Exec(ctx, insertTx,
		arg.TxHash,
		arg.EpochID,
		arg.TxBytes,
		arg.Status,
	)
	return err
}

const rejectNewTransactions = `-- name: RejectNewTransactions :exec
UPDATE transaction
SET status='rejected'
WHERE epoch_id=$1 AND status='new'
`

func (q *Queries) RejectNewTransactions(ctx context.Context, epochID []byte) error {
	_, err := q.db.Exec(ctx, rejectNewTransactions, epochID)
	return err
}

const setBatchSubmitted = `-- name: SetBatchSubmitted :exec
UPDATE batchtx SET submitted=true WHERE submitted=false
`

func (q *Queries) SetBatchSubmitted(ctx context.Context) error {
	_, err := q.db.Exec(ctx, setBatchSubmitted)
	return err
}

const setNextBatch = `-- name: SetNextBatch :exec
INSERT INTO next_batch (epoch_id, l1_block_number) VALUES ($1, $2)
ON CONFLICT (enforce_one_row) DO UPDATE
SET epoch_id = $1, l1_block_number = $2
`

type SetNextBatchParams struct {
	EpochID       []byte
	L1BlockNumber int64
}

func (q *Queries) SetNextBatch(ctx context.Context, arg SetNextBatchParams) error {
	_, err := q.db.Exec(ctx, setNextBatch, arg.EpochID, arg.L1BlockNumber)
	return err
}

const setTransactionStatus = `-- name: SetTransactionStatus :exec
UPDATE transaction
SET status=$2
WHERE tx_hash = $1
`

type SetTransactionStatusParams struct {
	TxHash []byte
	Status Txstatus
}

func (q *Queries) SetTransactionStatus(ctx context.Context, arg SetTransactionStatusParams) error {
	_, err := q.db.Exec(ctx, setTransactionStatus, arg.TxHash, arg.Status)
	return err
}

const updateDecryptionTriggerSent = `-- name: UpdateDecryptionTriggerSent :exec
UPDATE decryption_trigger
SET sent=NOW()
WHERE epoch_id=$1
`

func (q *Queries) UpdateDecryptionTriggerSent(ctx context.Context, epochID []byte) error {
	_, err := q.db.Exec(ctx, updateDecryptionTriggerSent, epochID)
	return err
}
