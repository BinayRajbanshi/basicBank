package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Store defines all functions to execute db queries and transactions
type Store interface {
	Querier // these are all methods under the hood
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
	// CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error)
	// VerifyEmailTx(ctx context.Context, arg VerifyEmailTxParams) (VerifyEmailTxResult, error)
}

// SQLStore provides all functions to execute SQL queries and transactions
// Important : SQLStore must also define the TransferTx and other transaction method to fully implement the Store interface. Hence they are defined in their respective files
type SQLStore struct {
	connPool *pgxpool.Pool
	*Queries
}

// NewStore creates a new store
func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}

// Why is returning &SQLStore{...} valid even though Store is an interface?
// This means:
// NewStore satisfies Store interface. This is because TransferTx function is extended with SQLStore so TransferTx is satisfied and also Queries has all the methods of querier.
