// Code generated by sqlc. DO NOT EDIT.

package kprdb

import (
	"database/sql"
	"time"
)

type ChainCollator struct {
	ActivationBlockNumber int64
	Collator              string
}

type DecryptionKey struct {
	EpochID       []byte
	DecryptionKey []byte
}

type DecryptionKeyShare struct {
	EpochID            []byte
	KeyperIndex        int64
	DecryptionKeyShare []byte
}

type DecryptionTrigger struct {
	EpochID []byte
}

type DkgResult struct {
	Eon        int64
	Success    bool
	Error      sql.NullString
	PureResult []byte
}

type Eon struct {
	Eon                   int64
	Height                int64
	ActivationBlockNumber int64
	ConfigIndex           int64
}

type EventSyncProgress struct {
	ID              bool
	NextBlockNumber int32
	NextLogIndex    int32
}

type KeyperSet struct {
	KeyperConfigIndex     int64
	ActivationBlockNumber int64
	Keypers               []string
	Threshold             int32
}

type LastBatchConfigSent struct {
	EnforceOneRow     bool
	KeyperConfigIndex int64
}

type LastBlockSeen struct {
	EnforceOneRow bool
	BlockNumber   int64
}

type MetaInf struct {
	Key   string
	Value string
}

type OutgoingEonKey struct {
	EonPublicKey []byte
	Eon          int64
}

type PolyEval struct {
	Eon             int64
	ReceiverAddress string
	Eval            []byte
}

type Puredkg struct {
	Eon     int64
	Puredkg []byte
}

type TendermintBatchConfig struct {
	ConfigIndex           int32
	Height                int64
	Keypers               []string
	Threshold             int32
	Started               bool
	ActivationBlockNumber int64
}

type TendermintEncryptionKey struct {
	Address             string
	EncryptionPublicKey []byte
}

type TendermintOutgoingMessage struct {
	ID          int32
	Description string
	Msg         []byte
}

type TendermintSyncMetum struct {
	CurrentBlock        int64
	LastCommittedHeight int64
	SyncTimestamp       time.Time
}
