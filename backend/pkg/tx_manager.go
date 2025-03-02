package pkg

import "context"

type Tx interface {
	Commit() error
	Rollback() error
}

// TxManager provides an abstraction for database transactions,
// allowing flexible transaction management.
type TxManager[T Tx] interface {
	// Begin starts a new transaction and returns a Tx instance
	Begin(ctx context.Context) (T, error)
}
