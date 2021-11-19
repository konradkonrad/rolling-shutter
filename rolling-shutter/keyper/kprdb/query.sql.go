// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package kprdb

import (
	"context"
	"database/sql"
	"time"

	"github.com/jackc/pgconn"
)

const countBatchConfigs = `-- name: CountBatchConfigs :one
SELECT count(*) FROM keyper.tendermint_batch_config
`

func (q *Queries) CountBatchConfigs(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, countBatchConfigs)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countDecryptionKeyShares = `-- name: CountDecryptionKeyShares :one
SELECT count(*) FROM keyper.decryption_key_share
WHERE epoch_id = $1
`

func (q *Queries) CountDecryptionKeyShares(ctx context.Context, epochID []byte) (int64, error) {
	row := q.db.QueryRow(ctx, countDecryptionKeyShares, epochID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const deletePolyEval = `-- name: DeletePolyEval :exec

DELETE FROM keyper.poly_evals ev WHERE ev.eon=$1 AND ev.receiver_address=$2
`

type DeletePolyEvalParams struct {
	Eon             int64
	ReceiverAddress string
}

// PolyEvalsWithEncryptionKeys could probably already delete the entries from the poly_evals table.
// I wasn't able to make this work, because of bugs in sqlc
func (q *Queries) DeletePolyEval(ctx context.Context, arg DeletePolyEvalParams) error {
	_, err := q.db.Exec(ctx, deletePolyEval, arg.Eon, arg.ReceiverAddress)
	return err
}

const deletePolyEvalByEon = `-- name: DeletePolyEvalByEon :execresult
DELETE FROM keyper.poly_evals ev WHERE ev.eon=$1
`

func (q *Queries) DeletePolyEvalByEon(ctx context.Context, eon int64) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, deletePolyEvalByEon, eon)
}

const deletePureDKG = `-- name: DeletePureDKG :exec
DELETE FROM keyper.puredkg WHERE eon=$1
`

func (q *Queries) DeletePureDKG(ctx context.Context, eon int64) error {
	_, err := q.db.Exec(ctx, deletePureDKG, eon)
	return err
}

const deleteShutterMessage = `-- name: DeleteShutterMessage :exec
DELETE FROM keyper.tendermint_outgoing_messages WHERE id=$1
`

func (q *Queries) DeleteShutterMessage(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteShutterMessage, id)
	return err
}

const existsDecryptionKey = `-- name: ExistsDecryptionKey :one
SELECT EXISTS (
    SELECT 1
    FROM keyper.decryption_key
    WHERE epoch_id = $1
)
`

func (q *Queries) ExistsDecryptionKey(ctx context.Context, epochID []byte) (bool, error) {
	row := q.db.QueryRow(ctx, existsDecryptionKey, epochID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const existsDecryptionKeyShare = `-- name: ExistsDecryptionKeyShare :one
SELECT EXISTS (
    SELECT 1
    FROM keyper.decryption_key_share
    WHERE epoch_id = $1 AND keyper_index = $2
)
`

type ExistsDecryptionKeyShareParams struct {
	EpochID     []byte
	KeyperIndex int64
}

func (q *Queries) ExistsDecryptionKeyShare(ctx context.Context, arg ExistsDecryptionKeyShareParams) (bool, error) {
	row := q.db.QueryRow(ctx, existsDecryptionKeyShare, arg.EpochID, arg.KeyperIndex)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const getBatchConfig = `-- name: GetBatchConfig :one
SELECT config_index, height, keypers, threshold
FROM keyper.tendermint_batch_config
WHERE config_index = $1
`

func (q *Queries) GetBatchConfig(ctx context.Context, configIndex int32) (KeyperTendermintBatchConfig, error) {
	row := q.db.QueryRow(ctx, getBatchConfig, configIndex)
	var i KeyperTendermintBatchConfig
	err := row.Scan(
		&i.ConfigIndex,
		&i.Height,
		&i.Keypers,
		&i.Threshold,
	)
	return i, err
}

const getBatchConfigs = `-- name: GetBatchConfigs :many
SELECT config_index, height, keypers, threshold
FROM keyper.tendermint_batch_config
ORDER BY config_index
`

func (q *Queries) GetBatchConfigs(ctx context.Context) ([]KeyperTendermintBatchConfig, error) {
	rows, err := q.db.Query(ctx, getBatchConfigs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []KeyperTendermintBatchConfig
	for rows.Next() {
		var i KeyperTendermintBatchConfig
		if err := rows.Scan(
			&i.ConfigIndex,
			&i.Height,
			&i.Keypers,
			&i.Threshold,
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

const getDKGResult = `-- name: GetDKGResult :one
SELECT eon, success, error, pure_result FROM keyper.dkg_result
WHERE eon = $1
`

func (q *Queries) GetDKGResult(ctx context.Context, eon int64) (KeyperDkgResult, error) {
	row := q.db.QueryRow(ctx, getDKGResult, eon)
	var i KeyperDkgResult
	err := row.Scan(
		&i.Eon,
		&i.Success,
		&i.Error,
		&i.PureResult,
	)
	return i, err
}

const getDKGResultForBlockNumber = `-- name: GetDKGResultForBlockNumber :one
SELECT eon, success, error, pure_result FROM keyper.dkg_result
WHERE eon = (SELECT eon FROM keyper.eons WHERE activation_block_number <= $1
ORDER BY activation_block_number DESC, height DESC
LIMIT 1)
`

func (q *Queries) GetDKGResultForBlockNumber(ctx context.Context, blockNumber int64) (KeyperDkgResult, error) {
	row := q.db.QueryRow(ctx, getDKGResultForBlockNumber, blockNumber)
	var i KeyperDkgResult
	err := row.Scan(
		&i.Eon,
		&i.Success,
		&i.Error,
		&i.PureResult,
	)
	return i, err
}

const getDecryptionKey = `-- name: GetDecryptionKey :one
SELECT epoch_id, decryption_key FROM keyper.decryption_key
WHERE epoch_id = $1
`

func (q *Queries) GetDecryptionKey(ctx context.Context, epochID []byte) (KeyperDecryptionKey, error) {
	row := q.db.QueryRow(ctx, getDecryptionKey, epochID)
	var i KeyperDecryptionKey
	err := row.Scan(&i.EpochID, &i.DecryptionKey)
	return i, err
}

const getDecryptionKeyShare = `-- name: GetDecryptionKeyShare :one
SELECT epoch_id, keyper_index, decryption_key_share FROM keyper.decryption_key_share
WHERE epoch_id = $1 AND keyper_index = $2
`

type GetDecryptionKeyShareParams struct {
	EpochID     []byte
	KeyperIndex int64
}

func (q *Queries) GetDecryptionKeyShare(ctx context.Context, arg GetDecryptionKeyShareParams) (KeyperDecryptionKeyShare, error) {
	row := q.db.QueryRow(ctx, getDecryptionKeyShare, arg.EpochID, arg.KeyperIndex)
	var i KeyperDecryptionKeyShare
	err := row.Scan(&i.EpochID, &i.KeyperIndex, &i.DecryptionKeyShare)
	return i, err
}

const getEncryptionKeys = `-- name: GetEncryptionKeys :many
SELECT address, encryption_public_key FROM keyper.tendermint_encryption_key
`

func (q *Queries) GetEncryptionKeys(ctx context.Context) ([]KeyperTendermintEncryptionKey, error) {
	rows, err := q.db.Query(ctx, getEncryptionKeys)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []KeyperTendermintEncryptionKey
	for rows.Next() {
		var i KeyperTendermintEncryptionKey
		if err := rows.Scan(&i.Address, &i.EncryptionPublicKey); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEon = `-- name: GetEon :one
SELECT eon, height, activation_block_number, config_index FROM keyper.eons WHERE eon=$1
`

func (q *Queries) GetEon(ctx context.Context, eon int64) (KeyperEon, error) {
	row := q.db.QueryRow(ctx, getEon, eon)
	var i KeyperEon
	err := row.Scan(
		&i.Eon,
		&i.Height,
		&i.ActivationBlockNumber,
		&i.ConfigIndex,
	)
	return i, err
}

const getEonForBlockNumber = `-- name: GetEonForBlockNumber :one
SELECT eon, height, activation_block_number, config_index FROM keyper.eons
WHERE activation_block_number <= $1
ORDER BY activation_block_number DESC, height DESC
LIMIT 1
`

func (q *Queries) GetEonForBlockNumber(ctx context.Context, blockNumber int64) (KeyperEon, error) {
	row := q.db.QueryRow(ctx, getEonForBlockNumber, blockNumber)
	var i KeyperEon
	err := row.Scan(
		&i.Eon,
		&i.Height,
		&i.ActivationBlockNumber,
		&i.ConfigIndex,
	)
	return i, err
}

const getEventSyncProgress = `-- name: GetEventSyncProgress :one
SELECT id, next_block_number, next_log_index FROM keyper.event_sync_progress LIMIT 1
`

func (q *Queries) GetEventSyncProgress(ctx context.Context) (KeyperEventSyncProgress, error) {
	row := q.db.QueryRow(ctx, getEventSyncProgress)
	var i KeyperEventSyncProgress
	err := row.Scan(&i.ID, &i.NextBlockNumber, &i.NextLogIndex)
	return i, err
}

const getLastCommittedHeight = `-- name: GetLastCommittedHeight :one
SELECT last_committed_height
FROM keyper.tendermint_sync_meta
ORDER BY current_block DESC, last_committed_height DESC
LIMIT 1
`

func (q *Queries) GetLastCommittedHeight(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, getLastCommittedHeight)
	var last_committed_height int64
	err := row.Scan(&last_committed_height)
	return last_committed_height, err
}

const getLatestBatchConfig = `-- name: GetLatestBatchConfig :one
SELECT config_index, height, keypers, threshold
FROM keyper.tendermint_batch_config
ORDER BY config_index DESC
LIMIT 1
`

func (q *Queries) GetLatestBatchConfig(ctx context.Context) (KeyperTendermintBatchConfig, error) {
	row := q.db.QueryRow(ctx, getLatestBatchConfig)
	var i KeyperTendermintBatchConfig
	err := row.Scan(
		&i.ConfigIndex,
		&i.Height,
		&i.Keypers,
		&i.Threshold,
	)
	return i, err
}

const getMeta = `-- name: GetMeta :one
SELECT value FROM keyper.meta_inf WHERE key = $1
`

func (q *Queries) GetMeta(ctx context.Context, key string) (string, error) {
	row := q.db.QueryRow(ctx, getMeta, key)
	var value string
	err := row.Scan(&value)
	return value, err
}

const getNextShutterMessage = `-- name: GetNextShutterMessage :one
SELECT id, description, msg from keyper.tendermint_outgoing_messages
ORDER BY id
LIMIT 1
`

func (q *Queries) GetNextShutterMessage(ctx context.Context) (KeyperTendermintOutgoingMessage, error) {
	row := q.db.QueryRow(ctx, getNextShutterMessage)
	var i KeyperTendermintOutgoingMessage
	err := row.Scan(&i.ID, &i.Description, &i.Msg)
	return i, err
}

const insertBatchConfig = `-- name: InsertBatchConfig :exec
INSERT INTO keyper.tendermint_batch_config (config_index, height, keypers, threshold)
VALUES ($1, $2, $3, $4)
`

type InsertBatchConfigParams struct {
	ConfigIndex int32
	Height      int64
	Keypers     []string
	Threshold   int32
}

func (q *Queries) InsertBatchConfig(ctx context.Context, arg InsertBatchConfigParams) error {
	_, err := q.db.Exec(ctx, insertBatchConfig,
		arg.ConfigIndex,
		arg.Height,
		arg.Keypers,
		arg.Threshold,
	)
	return err
}

const insertDKGResult = `-- name: InsertDKGResult :exec
INSERT INTO keyper.dkg_result (eon,success,error,pure_result)
VALUES ($1,$2,$3,$4)
`

type InsertDKGResultParams struct {
	Eon        int64
	Success    bool
	Error      sql.NullString
	PureResult []byte
}

func (q *Queries) InsertDKGResult(ctx context.Context, arg InsertDKGResultParams) error {
	_, err := q.db.Exec(ctx, insertDKGResult,
		arg.Eon,
		arg.Success,
		arg.Error,
		arg.PureResult,
	)
	return err
}

const insertDecryptionKey = `-- name: InsertDecryptionKey :execresult
INSERT INTO keyper.decryption_key (epoch_id, decryption_key)
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

const insertDecryptionKeyShare = `-- name: InsertDecryptionKeyShare :exec
INSERT INTO keyper.decryption_key_share (epoch_id, keyper_index, decryption_key_share)
VALUES ($1, $2, $3)
`

type InsertDecryptionKeyShareParams struct {
	EpochID            []byte
	KeyperIndex        int64
	DecryptionKeyShare []byte
}

func (q *Queries) InsertDecryptionKeyShare(ctx context.Context, arg InsertDecryptionKeyShareParams) error {
	_, err := q.db.Exec(ctx, insertDecryptionKeyShare, arg.EpochID, arg.KeyperIndex, arg.DecryptionKeyShare)
	return err
}

const insertEncryptionKey = `-- name: InsertEncryptionKey :exec
INSERT INTO keyper.tendermint_encryption_key (address, encryption_public_key) VALUES ($1, $2)
`

type InsertEncryptionKeyParams struct {
	Address             string
	EncryptionPublicKey []byte
}

func (q *Queries) InsertEncryptionKey(ctx context.Context, arg InsertEncryptionKeyParams) error {
	_, err := q.db.Exec(ctx, insertEncryptionKey, arg.Address, arg.EncryptionPublicKey)
	return err
}

const insertEon = `-- name: InsertEon :exec
INSERT INTO keyper.eons (eon, height, activation_block_number, config_index)
VALUES ($1, $2, $3, $4)
`

type InsertEonParams struct {
	Eon                   int64
	Height                int64
	ActivationBlockNumber int64
	ConfigIndex           int64
}

func (q *Queries) InsertEon(ctx context.Context, arg InsertEonParams) error {
	_, err := q.db.Exec(ctx, insertEon,
		arg.Eon,
		arg.Height,
		arg.ActivationBlockNumber,
		arg.ConfigIndex,
	)
	return err
}

const insertMeta = `-- name: InsertMeta :exec
INSERT INTO keyper.meta_inf (key, value) VALUES ($1, $2)
`

type InsertMetaParams struct {
	Key   string
	Value string
}

func (q *Queries) InsertMeta(ctx context.Context, arg InsertMetaParams) error {
	_, err := q.db.Exec(ctx, insertMeta, arg.Key, arg.Value)
	return err
}

const insertPolyEval = `-- name: InsertPolyEval :exec
INSERT INTO keyper.poly_evals (eon, receiver_address, eval)
VALUES ($1, $2, $3)
`

type InsertPolyEvalParams struct {
	Eon             int64
	ReceiverAddress string
	Eval            []byte
}

func (q *Queries) InsertPolyEval(ctx context.Context, arg InsertPolyEvalParams) error {
	_, err := q.db.Exec(ctx, insertPolyEval, arg.Eon, arg.ReceiverAddress, arg.Eval)
	return err
}

const insertPureDKG = `-- name: InsertPureDKG :exec
INSERT INTO keyper.puredkg (eon, puredkg) VALUES ($1, $2)
ON CONFLICT (eon) DO UPDATE SET puredkg=EXCLUDED.puredkg
`

type InsertPureDKGParams struct {
	Eon     int64
	Puredkg []byte
}

func (q *Queries) InsertPureDKG(ctx context.Context, arg InsertPureDKGParams) error {
	_, err := q.db.Exec(ctx, insertPureDKG, arg.Eon, arg.Puredkg)
	return err
}

const polyEvalsWithEncryptionKeys = `-- name: PolyEvalsWithEncryptionKeys :many
SELECT ev.eon, ev.receiver_address, ev.eval,
       k.encryption_public_key,
       eons.height
FROM keyper.poly_evals ev
INNER JOIN keyper.tendermint_encryption_key k
      ON ev.receiver_address = k.address
INNER JOIN keyper.eons eons
      ON ev.eon = eons.eon
ORDER BY ev.eon
`

type PolyEvalsWithEncryptionKeysRow struct {
	Eon                 int64
	ReceiverAddress     string
	Eval                []byte
	EncryptionPublicKey []byte
	Height              int64
}

func (q *Queries) PolyEvalsWithEncryptionKeys(ctx context.Context) ([]PolyEvalsWithEncryptionKeysRow, error) {
	rows, err := q.db.Query(ctx, polyEvalsWithEncryptionKeys)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PolyEvalsWithEncryptionKeysRow
	for rows.Next() {
		var i PolyEvalsWithEncryptionKeysRow
		if err := rows.Scan(
			&i.Eon,
			&i.ReceiverAddress,
			&i.Eval,
			&i.EncryptionPublicKey,
			&i.Height,
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

const scheduleShutterMessage = `-- name: ScheduleShutterMessage :one
INSERT INTO keyper.tendermint_outgoing_messages (description, msg)
VALUES ($1, $2)
RETURNING id
`

type ScheduleShutterMessageParams struct {
	Description string
	Msg         []byte
}

func (q *Queries) ScheduleShutterMessage(ctx context.Context, arg ScheduleShutterMessageParams) (int32, error) {
	row := q.db.QueryRow(ctx, scheduleShutterMessage, arg.Description, arg.Msg)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const selectDecryptionKeyShares = `-- name: SelectDecryptionKeyShares :many
SELECT epoch_id, keyper_index, decryption_key_share FROM keyper.decryption_key_share
WHERE epoch_id = $1
`

func (q *Queries) SelectDecryptionKeyShares(ctx context.Context, epochID []byte) ([]KeyperDecryptionKeyShare, error) {
	rows, err := q.db.Query(ctx, selectDecryptionKeyShares, epochID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []KeyperDecryptionKeyShare
	for rows.Next() {
		var i KeyperDecryptionKeyShare
		if err := rows.Scan(&i.EpochID, &i.KeyperIndex, &i.DecryptionKeyShare); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectPureDKG = `-- name: SelectPureDKG :many
SELECT eon, puredkg FROM keyper.puredkg
`

func (q *Queries) SelectPureDKG(ctx context.Context) ([]KeyperPuredkg, error) {
	rows, err := q.db.Query(ctx, selectPureDKG)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []KeyperPuredkg
	for rows.Next() {
		var i KeyperPuredkg
		if err := rows.Scan(&i.Eon, &i.Puredkg); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const tMGetSyncMeta = `-- name: TMGetSyncMeta :one
SELECT current_block, last_committed_height, sync_timestamp
FROM keyper.tendermint_sync_meta
ORDER BY current_block DESC, last_committed_height DESC
LIMIT 1
`

func (q *Queries) TMGetSyncMeta(ctx context.Context) (KeyperTendermintSyncMetum, error) {
	row := q.db.QueryRow(ctx, tMGetSyncMeta)
	var i KeyperTendermintSyncMetum
	err := row.Scan(&i.CurrentBlock, &i.LastCommittedHeight, &i.SyncTimestamp)
	return i, err
}

const tMSetSyncMeta = `-- name: TMSetSyncMeta :exec
INSERT INTO keyper.tendermint_sync_meta (current_block, last_committed_height, sync_timestamp)
VALUES ($1, $2, $3)
`

type TMSetSyncMetaParams struct {
	CurrentBlock        int64
	LastCommittedHeight int64
	SyncTimestamp       time.Time
}

func (q *Queries) TMSetSyncMeta(ctx context.Context, arg TMSetSyncMetaParams) error {
	_, err := q.db.Exec(ctx, tMSetSyncMeta, arg.CurrentBlock, arg.LastCommittedHeight, arg.SyncTimestamp)
	return err
}

const updateEventSyncProgress = `-- name: UpdateEventSyncProgress :exec
INSERT INTO keyper.event_sync_progress (next_block_number, next_log_index)
VALUES ($1, $2)
ON CONFLICT (id) DO UPDATE
    SET next_block_number = $1,
        next_log_index = $2
`

type UpdateEventSyncProgressParams struct {
	NextBlockNumber int32
	NextLogIndex    int32
}

func (q *Queries) UpdateEventSyncProgress(ctx context.Context, arg UpdateEventSyncProgressParams) error {
	_, err := q.db.Exec(ctx, updateEventSyncProgress, arg.NextBlockNumber, arg.NextLogIndex)
	return err
}
