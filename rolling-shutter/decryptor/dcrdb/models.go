// Code generated by sqlc. DO NOT EDIT.

package dcrdb

import ()

type DecryptorCipherBatch struct {
	EpochID []byte
	Data    []byte
}

type DecryptorDecryptionKey struct {
	EpochID []byte
	Key     []byte
}

type DecryptorDecryptionSignature struct {
	EpochID     []byte
	SignedHash  []byte
	SignerIndex int64
	Signature   []byte
}

type DecryptorMetaInf struct {
	Key   string
	Value string
}
