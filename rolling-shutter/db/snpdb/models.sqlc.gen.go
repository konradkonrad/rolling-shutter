// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package snpdb

import ()

type DecryptionKey struct {
	EpochID []byte
	Key     []byte
}

type EonPublicKey struct {
	EonID        int64
	EonPublicKey []byte
}