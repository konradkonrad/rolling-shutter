// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package chainobsdb

import ()

type ChainCollator struct {
	ActivationBlockNumber int64
	Collator              string
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
