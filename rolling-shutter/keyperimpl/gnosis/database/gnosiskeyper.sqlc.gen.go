// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: gnosiskeyper.sql

package database

import (
	"context"

	"github.com/jackc/pgconn"
)

const getCurrentDecryptionTrigger = `-- name: GetCurrentDecryptionTrigger :one
SELECT eon, slot, tx_pointer, identities_hash FROM current_decryption_trigger
WHERE eon = $1
`

func (q *Queries) GetCurrentDecryptionTrigger(ctx context.Context, eon int64) (CurrentDecryptionTrigger, error) {
	row := q.db.QueryRow(ctx, getCurrentDecryptionTrigger, eon)
	var i CurrentDecryptionTrigger
	err := row.Scan(
		&i.Eon,
		&i.Slot,
		&i.TxPointer,
		&i.IdentitiesHash,
	)
	return i, err
}

const getSlotDecryptionSignatures = `-- name: GetSlotDecryptionSignatures :many
SELECT eon, slot, keyper_index, tx_pointer, identities_hash, signature FROM slot_decryption_signatures
WHERE eon = $1 AND slot = $2 AND tx_pointer = $3 AND identities_hash = $4
ORDER BY keyper_index ASC
LIMIT $5
`

type GetSlotDecryptionSignaturesParams struct {
	Eon            int64
	Slot           int64
	TxPointer      int64
	IdentitiesHash []byte
	Limit          int32
}

func (q *Queries) GetSlotDecryptionSignatures(ctx context.Context, arg GetSlotDecryptionSignaturesParams) ([]SlotDecryptionSignature, error) {
	rows, err := q.db.Query(ctx, getSlotDecryptionSignatures,
		arg.Eon,
		arg.Slot,
		arg.TxPointer,
		arg.IdentitiesHash,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SlotDecryptionSignature
	for rows.Next() {
		var i SlotDecryptionSignature
		if err := rows.Scan(
			&i.Eon,
			&i.Slot,
			&i.KeyperIndex,
			&i.TxPointer,
			&i.IdentitiesHash,
			&i.Signature,
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

const getTransactionSubmittedEventCount = `-- name: GetTransactionSubmittedEventCount :one
SELECT event_count FROM transaction_submitted_event_count
WHERE eon = $1
LIMIT 1
`

func (q *Queries) GetTransactionSubmittedEventCount(ctx context.Context, eon int64) (int64, error) {
	row := q.db.QueryRow(ctx, getTransactionSubmittedEventCount, eon)
	var event_count int64
	err := row.Scan(&event_count)
	return event_count, err
}

const getTransactionSubmittedEvents = `-- name: GetTransactionSubmittedEvents :many
SELECT index, block_number, block_hash, tx_index, log_index, eon, identity_prefix, sender, gas_limit FROM transaction_submitted_event
WHERE eon = $1 AND index >= $2
ORDER BY index ASC
LIMIT $3
`

type GetTransactionSubmittedEventsParams struct {
	Eon   int64
	Index int64
	Limit int32
}

func (q *Queries) GetTransactionSubmittedEvents(ctx context.Context, arg GetTransactionSubmittedEventsParams) ([]TransactionSubmittedEvent, error) {
	rows, err := q.db.Query(ctx, getTransactionSubmittedEvents, arg.Eon, arg.Index, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TransactionSubmittedEvent
	for rows.Next() {
		var i TransactionSubmittedEvent
		if err := rows.Scan(
			&i.Index,
			&i.BlockNumber,
			&i.BlockHash,
			&i.TxIndex,
			&i.LogIndex,
			&i.Eon,
			&i.IdentityPrefix,
			&i.Sender,
			&i.GasLimit,
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

const getTransactionSubmittedEventsSyncedUntil = `-- name: GetTransactionSubmittedEventsSyncedUntil :one
SELECT enforce_one_row, block_hash, block_number, slot FROM transaction_submitted_events_synced_until LIMIT 1
`

func (q *Queries) GetTransactionSubmittedEventsSyncedUntil(ctx context.Context) (TransactionSubmittedEventsSyncedUntil, error) {
	row := q.db.QueryRow(ctx, getTransactionSubmittedEventsSyncedUntil)
	var i TransactionSubmittedEventsSyncedUntil
	err := row.Scan(
		&i.EnforceOneRow,
		&i.BlockHash,
		&i.BlockNumber,
		&i.Slot,
	)
	return i, err
}

const getTxPointer = `-- name: GetTxPointer :one
SELECT eon, slot, value FROM tx_pointer
WHERE eon = $1
`

func (q *Queries) GetTxPointer(ctx context.Context, eon int64) (TxPointer, error) {
	row := q.db.QueryRow(ctx, getTxPointer, eon)
	var i TxPointer
	err := row.Scan(&i.Eon, &i.Slot, &i.Value)
	return i, err
}

const getValidatorRegistrationNonceBefore = `-- name: GetValidatorRegistrationNonceBefore :one
SELECT nonce FROM validator_registrations
WHERE validator_index = $1 AND block_number <= $2 AND tx_index <= $3 AND log_index <= $4
ORDER BY block_number DESC, tx_index DESC, log_index DESC
LIMIT 1
`

type GetValidatorRegistrationNonceBeforeParams struct {
	ValidatorIndex int64
	BlockNumber    int64
	TxIndex        int64
	LogIndex       int64
}

func (q *Queries) GetValidatorRegistrationNonceBefore(ctx context.Context, arg GetValidatorRegistrationNonceBeforeParams) (int64, error) {
	row := q.db.QueryRow(ctx, getValidatorRegistrationNonceBefore,
		arg.ValidatorIndex,
		arg.BlockNumber,
		arg.TxIndex,
		arg.LogIndex,
	)
	var nonce int64
	err := row.Scan(&nonce)
	return nonce, err
}

const getValidatorRegistrationsSyncedUntil = `-- name: GetValidatorRegistrationsSyncedUntil :one
SELECT enforce_one_row, block_hash, block_number FROM validator_registrations_synced_until LIMIT 1
`

func (q *Queries) GetValidatorRegistrationsSyncedUntil(ctx context.Context) (ValidatorRegistrationsSyncedUntil, error) {
	row := q.db.QueryRow(ctx, getValidatorRegistrationsSyncedUntil)
	var i ValidatorRegistrationsSyncedUntil
	err := row.Scan(&i.EnforceOneRow, &i.BlockHash, &i.BlockNumber)
	return i, err
}

const initTxPointer = `-- name: InitTxPointer :exec
INSERT INTO tx_pointer (eon, slot, value)
VALUES ($1, $2, 0)
ON CONFLICT DO NOTHING
`

type InitTxPointerParams struct {
	Eon  int64
	Slot int64
}

func (q *Queries) InitTxPointer(ctx context.Context, arg InitTxPointerParams) error {
	_, err := q.db.Exec(ctx, initTxPointer, arg.Eon, arg.Slot)
	return err
}

const insertSlotDecryptionSignature = `-- name: InsertSlotDecryptionSignature :exec
INSERT INTO slot_decryption_signatures (eon, slot, keyper_index, tx_pointer, identities_hash, signature)
VALUES ($1, $2, $3, $4, $5, $6)
`

type InsertSlotDecryptionSignatureParams struct {
	Eon            int64
	Slot           int64
	KeyperIndex    int64
	TxPointer      int64
	IdentitiesHash []byte
	Signature      []byte
}

func (q *Queries) InsertSlotDecryptionSignature(ctx context.Context, arg InsertSlotDecryptionSignatureParams) error {
	_, err := q.db.Exec(ctx, insertSlotDecryptionSignature,
		arg.Eon,
		arg.Slot,
		arg.KeyperIndex,
		arg.TxPointer,
		arg.IdentitiesHash,
		arg.Signature,
	)
	return err
}

const insertTransactionSubmittedEvent = `-- name: InsertTransactionSubmittedEvent :execresult
INSERT INTO transaction_submitted_event (
    index,
    block_number,
    block_hash,
    tx_index,
    log_index,
    eon,
    identity_prefix,
    sender,
    gas_limit
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
ON CONFLICT DO NOTHING
`

type InsertTransactionSubmittedEventParams struct {
	Index          int64
	BlockNumber    int64
	BlockHash      []byte
	TxIndex        int64
	LogIndex       int64
	Eon            int64
	IdentityPrefix []byte
	Sender         string
	GasLimit       int64
}

func (q *Queries) InsertTransactionSubmittedEvent(ctx context.Context, arg InsertTransactionSubmittedEventParams) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, insertTransactionSubmittedEvent,
		arg.Index,
		arg.BlockNumber,
		arg.BlockHash,
		arg.TxIndex,
		arg.LogIndex,
		arg.Eon,
		arg.IdentityPrefix,
		arg.Sender,
		arg.GasLimit,
	)
}

const insertValidatorRegistration = `-- name: InsertValidatorRegistration :exec
INSERT INTO validator_registrations (
    block_number,
    block_hash,
    tx_index,
    log_index,
    validator_index,
    nonce,
    is_registration
) VALUES ($1, $2, $3, $4, $5, $6, $7)
`

type InsertValidatorRegistrationParams struct {
	BlockNumber    int64
	BlockHash      []byte
	TxIndex        int64
	LogIndex       int64
	ValidatorIndex int64
	Nonce          int64
	IsRegistration bool
}

func (q *Queries) InsertValidatorRegistration(ctx context.Context, arg InsertValidatorRegistrationParams) error {
	_, err := q.db.Exec(ctx, insertValidatorRegistration,
		arg.BlockNumber,
		arg.BlockHash,
		arg.TxIndex,
		arg.LogIndex,
		arg.ValidatorIndex,
		arg.Nonce,
		arg.IsRegistration,
	)
	return err
}

const isValidatorRegistered = `-- name: IsValidatorRegistered :one
SELECT is_registration FROM validator_registrations
WHERE validator_index = $1 AND block_number < $2
ORDER BY block_number DESC, tx_index DESC, log_index DESC
LIMIT 1
`

type IsValidatorRegisteredParams struct {
	ValidatorIndex int64
	BlockNumber    int64
}

func (q *Queries) IsValidatorRegistered(ctx context.Context, arg IsValidatorRegisteredParams) (bool, error) {
	row := q.db.QueryRow(ctx, isValidatorRegistered, arg.ValidatorIndex, arg.BlockNumber)
	var is_registration bool
	err := row.Scan(&is_registration)
	return is_registration, err
}

const setCurrentDecryptionTrigger = `-- name: SetCurrentDecryptionTrigger :exec
INSERT INTO current_decryption_trigger (eon, slot, tx_pointer, identities_hash)
VALUES ($1, $2, $3, $4)
ON CONFLICT (eon) DO UPDATE
SET slot = $2, tx_pointer = $3, identities_hash = $4
`

type SetCurrentDecryptionTriggerParams struct {
	Eon            int64
	Slot           int64
	TxPointer      int64
	IdentitiesHash []byte
}

func (q *Queries) SetCurrentDecryptionTrigger(ctx context.Context, arg SetCurrentDecryptionTriggerParams) error {
	_, err := q.db.Exec(ctx, setCurrentDecryptionTrigger,
		arg.Eon,
		arg.Slot,
		arg.TxPointer,
		arg.IdentitiesHash,
	)
	return err
}

const setTransactionSubmittedEventCount = `-- name: SetTransactionSubmittedEventCount :exec
INSERT INTO transaction_submitted_event_count (eon, event_count)
VALUES ($1, $2)
ON CONFLICT (eon) DO UPDATE
SET event_count = $2
`

type SetTransactionSubmittedEventCountParams struct {
	Eon        int64
	EventCount int64
}

func (q *Queries) SetTransactionSubmittedEventCount(ctx context.Context, arg SetTransactionSubmittedEventCountParams) error {
	_, err := q.db.Exec(ctx, setTransactionSubmittedEventCount, arg.Eon, arg.EventCount)
	return err
}

const setTransactionSubmittedEventsSyncedUntil = `-- name: SetTransactionSubmittedEventsSyncedUntil :exec
INSERT INTO transaction_submitted_events_synced_until (block_hash, block_number, slot) VALUES ($1, $2, $3)
ON CONFLICT (enforce_one_row) DO UPDATE
SET block_hash = $1, block_number = $2, slot = $3
`

type SetTransactionSubmittedEventsSyncedUntilParams struct {
	BlockHash   []byte
	BlockNumber int64
	Slot        int64
}

func (q *Queries) SetTransactionSubmittedEventsSyncedUntil(ctx context.Context, arg SetTransactionSubmittedEventsSyncedUntilParams) error {
	_, err := q.db.Exec(ctx, setTransactionSubmittedEventsSyncedUntil, arg.BlockHash, arg.BlockNumber, arg.Slot)
	return err
}

const setTxPointer = `-- name: SetTxPointer :exec
INSERT INTO tx_pointer (eon, slot, value)
VALUES ($1, $2, $3)
ON CONFLICT (eon) DO UPDATE
SET slot = $2, value = $3
`

type SetTxPointerParams struct {
	Eon   int64
	Slot  int64
	Value int64
}

func (q *Queries) SetTxPointer(ctx context.Context, arg SetTxPointerParams) error {
	_, err := q.db.Exec(ctx, setTxPointer, arg.Eon, arg.Slot, arg.Value)
	return err
}

const setTxPointerSlot = `-- name: SetTxPointerSlot :exec
UPDATE tx_pointer
SET slot = $2
WHERE eon = $1
`

type SetTxPointerSlotParams struct {
	Eon  int64
	Slot int64
}

func (q *Queries) SetTxPointerSlot(ctx context.Context, arg SetTxPointerSlotParams) error {
	_, err := q.db.Exec(ctx, setTxPointerSlot, arg.Eon, arg.Slot)
	return err
}

const setValidatorRegistrationsSyncedUntil = `-- name: SetValidatorRegistrationsSyncedUntil :exec
INSERT INTO validator_registrations_synced_until (block_hash, block_number) VALUES ($1, $2)
ON CONFLICT (enforce_one_row) DO UPDATE
SET block_hash = $1, block_number = $2
`

type SetValidatorRegistrationsSyncedUntilParams struct {
	BlockHash   []byte
	BlockNumber int64
}

func (q *Queries) SetValidatorRegistrationsSyncedUntil(ctx context.Context, arg SetValidatorRegistrationsSyncedUntilParams) error {
	_, err := q.db.Exec(ctx, setValidatorRegistrationsSyncedUntil, arg.BlockHash, arg.BlockNumber)
	return err
}
