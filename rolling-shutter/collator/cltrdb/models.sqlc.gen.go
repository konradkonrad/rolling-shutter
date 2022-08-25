// Code generated by sqlc. DO NOT EDIT.

package cltrdb

import ()

type DecryptionTrigger struct {
	EpochID   []byte
	BatchHash []byte
}

type MetaInf struct {
	Key   string
	Value string
}

type NextEpoch struct {
	EnforceOneRow bool
	EpochID       []byte
	BlockNumber   int64
}

type Transaction struct {
	TxID        []byte
	EpochID     []byte
	EncryptedTx []byte
}
