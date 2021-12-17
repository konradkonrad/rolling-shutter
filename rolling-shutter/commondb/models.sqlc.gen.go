// Code generated by sqlc. DO NOT EDIT.

package commondb

import ()

type ChainCollator struct {
	ActivationBlockNumber int64
	Collator              string
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
