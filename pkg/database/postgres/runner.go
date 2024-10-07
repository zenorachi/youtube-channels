package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Runner is an implementation of database connection with method WithTransaction.
// Exec - implementation of IExecutor interface.
type Runner struct {
	db   *sqlx.DB
	Exec IExecutor
}

// NewRunner creates a new Runner.
func NewRunner(db *sqlx.DB) *Runner {
	return &Runner{
		db:   db,
		Exec: db,
	}
}

// WithTransaction method is used when we need to call repository methods in a transaction.
// This method creates transaction and set it to Exec.
// When method is done, it sets Exec to default *sqlx.DB.
func (r *Runner) WithTransaction(work func() error) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	r.setTx(tx)
	defer r.unsetTx()

	if err = work(); err != nil {
		return fmt.Errorf(
			"work function error: %w | rollback error: %v",
			err,
			tx.Rollback(),
		)
	}

	return tx.Commit()
}

func (r *Runner) setTx(tx *sqlx.Tx) {
	r.Exec = tx
}

func (r *Runner) unsetTx() {
	r.Exec = r.db
}
