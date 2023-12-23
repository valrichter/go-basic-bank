package db

import (
	"context"
)

// Contains the imput parameters for CreateUser
type CreateUserTxParams struct {
	CreateUserParams
	AfterCreate func(user User) error
}

// Contains the result of a CreateUser transaction
type CreateUserTxResult struct {
	User User
}

// Money transfer from one account to another
// It create a transfer record, add account entries, and update account balances
// within a single database transaction.
func (store *SQLStore) CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error) {
	var result CreateUserTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.User, err = q.CreateUser(ctx, arg.CreateUserParams)
		if err != nil {
			return err
		}

		return arg.AfterCreate(result.User)
	})

	return result, err
}
