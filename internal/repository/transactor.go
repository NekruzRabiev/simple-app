package repository

import (
	"context"
	"fmt"
)

type transactorPostgres struct {
	db *store
}

func newTransactorPostgres(db *store) *transactorPostgres {
	return &transactorPostgres{db: db}
}

// WithinTransaction runs function within transaction
//
// The transaction commits when function were finished without error
func (r *transactorPostgres) WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error {
	// begin transaction
	tx, err := r.db.Beginx()
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	defer tx.Rollback()

	err = tFunc(injectTx(ctx, tx))
	if err != nil {
		return err
	}

	// if no error, commit
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
