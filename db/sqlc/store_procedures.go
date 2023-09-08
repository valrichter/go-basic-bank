package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute db queries & transactions
type Store struct {
	*Queries
	db *sql.DB
}

// Creates a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
		Queries: New(db),
	}
}

// Executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error { 
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// Contains the imput parameters for creating a new transfer
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID int64 `json:"to_account_id"`
	Amount float32 `json:"amount"`
}

// Contains the result of a transfer transaction
type TransferTxResult struct {
	Transfer Transfer `json:"transfer"`
	FromAccount Account `json:"from_account"`
	ToAccount Account `json:"to_account"`
	FromEntry Entry `json:"from_entry"`
	ToEntry Entry `json:"to_entry"`
}

// Money transfer from one account to another
// It create a transfer record, add account entries, and update account balances
// within a single database transaction.
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// 1. Make a transfer
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams {
			FromAccountID: arg.FromAccountID,
			ToAccountID: arg.ToAccountID,
			Amount: arg.Amount,
		})
		if err != nil {
			return err
		}
		

		// 2. Money is moving out of the account 'FromAccountID'
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams {
			AccountID: arg.FromAccountID,
			Amount: -arg.Amount,
		})
		if err != nil {
			return err
		}

		// 3. Money is moving into the account 'ToAccountID'
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams {
			AccountID: arg.ToAccountID,
			Amount: arg.Amount,
		})

		// TODO: 4. Update the accounts
		return nil
 	})
	return result, err
}