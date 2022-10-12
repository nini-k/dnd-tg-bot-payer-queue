package storage

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

type storage struct {
	db *sql.DB
}

func New(filePath string) (*storage, error) {
	db, err := sql.Open("sqlite3", filePath)
	if err != nil {
		return nil, errors.Wrap(err, "connecting")
	}

	return &storage{db: db}, nil
}

func (s *storage) GetSortedUsersByCurrentPayer(ctx context.Context) ([]User, error) {
	query := `
select 
    id,
    username,
    name,
    created_at
from user
`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "query processing")
	}
	defer rows.Close()

	out := make([]User, 0)
	for rows.Next() {
		user := User{}
		if err := rows.Scan(&user.Id, &user.Username, &user.Name, &user.CreatedAt); err != nil {
			return nil, errors.Wrap(err, "scanning rows")
		}

		out = append(out, user)
	}

	return out, nil
}

func (s *storage) ChangeCurrentPayer(ctx context.Context, nextUserId int, action UserPaymentHistoryAction) error {
	return s.execTransaction(ctx, func(txCtx context.Context, tx *sql.Tx) error {
		oldUserId, err := s.resetCurrentPayerTx(txCtx, tx)
		if err != nil {
			return errors.Wrapf(err, "reset current payer")
		}

		if err := s.setCurrentPayerTx(txCtx, tx, nextUserId, true); err != nil {
			return errors.Wrapf(err, "set next payer user_id=%d", nextUserId)
		}

		if err := s.addUserPaymentHistoryRecordTx(ctx, tx, oldUserId, action); err != nil {
			return errors.Wrapf(err, "add user payment history record user_id=%d action=%s", oldUserId, action)
		}

		return nil
	})
}

func (s *storage) addUserPaymentHistoryRecordTx(ctx context.Context, tx *sql.Tx, userId int64, action UserPaymentHistoryAction) error {
	query := `
insert into user_payment_history (user_id, action)
values (?, ?)
`

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}

	if _, err = stmt.ExecContext(ctx, userId, action); err != nil {
		return errors.Wrap(err, "query execution")
	}

	return nil
}

func (s *storage) resetCurrentPayerTx(ctx context.Context, tx *sql.Tx) (int64, error) {
	query := `
update user
set is_current_payer = false
where is_current_payer = true
returning id
`
	res, err := tx.ExecContext(ctx, query)
	if err != nil {
		return 0, errors.Wrap(err, "query execution")
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, errors.Wrap(err, "get updated user_id")
	}

	return id, nil
}

func (s *storage) setCurrentPayerTx(ctx context.Context, tx *sql.Tx, userId int, isCurrentPayer bool) error {
	query := `
update user
set is_current_payer = ?
where id = ?
`

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return errors.Wrap(err, "prepare statement")
	}

	if _, err = stmt.ExecContext(ctx, isCurrentPayer, userId); err != nil {
		return errors.Wrap(err, "query execution")
	}

	return nil
}

func (s *storage) execTransaction(ctx context.Context, fn func(context.Context, *sql.Tx) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "begin transaction")
	}

	if err = fn(ctx, tx); err != nil {
		if err = tx.Rollback(); err != nil {
			return errors.Wrap(err, "rollback transaction")
		}

		return err
	}

	if err = tx.Commit(); err != nil {
		return errors.Wrap(err, "commit transaction")
	}

	return nil
}
