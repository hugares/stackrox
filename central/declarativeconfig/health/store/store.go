package store

import (
	"context"

	"github.com/stackrox/rox/generated/storage"
)

// Store is the interface to the config health data layer
type Store interface {
	Get(ctx context.Context, id string) (*storage.DeclarativeConfigHealth, bool, error)
	Upsert(ctx context.Context, obj *storage.DeclarativeConfigHealth) error
	Delete(ctx context.Context, id string) error
	Walk(ctx context.Context, fn func(obj *storage.DeclarativeConfigHealth) error) error
}
