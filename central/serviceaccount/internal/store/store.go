// Code generated by boltbindings generator. DO NOT EDIT.

package store

import (
	bbolt "github.com/etcd-io/bbolt"
	storage "github.com/stackrox/rox/generated/storage"
)

type Store interface {
	DeleteServiceAccount(id string) error
	GetServiceAccount(id string) (*storage.ServiceAccount, bool, error)
	GetServiceAccounts(ids []string) ([]*storage.ServiceAccount, []int, error)
	ListServiceAccounts() ([]*storage.ServiceAccount, error)
	UpsertServiceAccount(serviceaccount *storage.ServiceAccount) error
}

func New(db *bbolt.DB) (Store, error) {
	return newStore(db)
}
