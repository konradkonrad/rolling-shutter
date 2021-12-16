// Code generated by sqlc. DO NOT EDIT.

package dcrdb

import ()

type AggregatedSignature struct {
	EpochID         []byte
	SignedHash      []byte
	SignersBitfield []byte
	Signature       []byte
}

type ChainCollator struct {
	ActivationBlockNumber int64
	Collator              string
}

type CipherBatch struct {
	EpochID      []byte
	Transactions [][]byte
}

type DecryptionKey struct {
	EpochID []byte
	Key     []byte
}

type DecryptionSignature struct {
	EpochID         []byte
	SignedHash      []byte
	SignersBitfield []byte
	Signature       []byte
}

type DecryptorIdentity struct {
	Address        string
	BlsPublicKey   []byte
	BlsSignature   []byte
	SignatureValid bool
}

type DecryptorSetMember struct {
	ActivationBlockNumber int64
	Index                 int32
	Address               string
}

type EonPublicKey struct {
	ActivationBlockNumber int64
	EonPublicKey          []byte
}

type EventSyncProgress struct {
	ID              bool
	NextBlockNumber int32
	NextLogIndex    int32
}

type KeyperSet struct {
	ActivationBlockNumber int64
	Keypers               []string
	Threshold             int32
}

type MetaInf struct {
	Key   string
	Value string
}
