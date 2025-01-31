package repository

import (
	"context"
	"errors"

	golog "github.com/Vladroon22/GoLog"
	"github.com/Vladroon22/TestTask-Bank-Operation/internal/entity"
	"github.com/jackc/pgx/v5"
	pool "github.com/jackc/pgx/v5/pgxpool"
)

type InterfaceRepo interface {
	IncreaseUserBalance(c context.Context, userID int, amount float64) error
	TransferMoney(c context.Context, userFrom, userTo string, fromUserID, toUserID int, amount float64) error
	GetLastTxs(c context.Context, userID int) ([]entity.Tx, error)
}

type Repository struct {
	db     *pool.Pool
	logger *golog.Logger
}

func NewRepository(db *pool.Pool, lg *golog.Logger) *Repository {
	return &Repository{db: db, logger: lg}
}

func (r *Repository) IncreaseUserBalance(c context.Context, userID int, amount float64) error {
	ctx, cancel := context.WithCancel(c)
	defer cancel()

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		r.logger.Errorln("Beg Tx (increase balance): ", err)
		return err
	}

	defer func() {
		errRb := tx.Rollback(ctx)
		if errRb != nil && !errors.Is(errRb, pgx.ErrTxClosed) {
			r.logger.Errorln("Rollback Tx (increase balance): ", errRb)
		}
	}()

	if _, err := tx.Exec(ctx, "UPDATE users SET balance = balance + $1 WHERE id = $2", amount, userID); err != nil {
		r.logger.Errorln("Update user's balance: ", err)
		return err
	}

	if _, err := tx.Exec(ctx, "INSERT INTO tx (user_id, amount, type) VALUES ($1, $2, $3)", userID, amount, "top_up"); err != nil {
		r.logger.Errorln("Insert into tx: ", err)
		return err
	}

	errTx := tx.Commit(ctx)
	if errTx != nil {
		r.logger.Errorln("Commit tx's error: ", err)
		return err
	}
	r.logger.Infoln("IncreaseBalance: success")

	return nil
}

func (r *Repository) TransferMoney(c context.Context, userFrom, userTo string, fromUserID, toUserID int, amount float64) error {
	ctx, cancel := context.WithCancel(c)
	defer cancel()

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		r.logger.Errorln("Beg Tx (transfer): ", err)
		return err
	}

	defer func() {
		errRb := tx.Rollback(ctx)
		if errRb != nil && !errors.Is(errRb, pgx.ErrTxClosed) {
			r.logger.Errorln("Rollback Tx (transfer): ", errRb)
		}
	}()

	if _, err := tx.Exec(ctx, "UPDATE users SET balance = balance - $1 WHERE id = $2", amount, fromUserID); err != nil {
		r.logger.Errorln("Update user's -->  balance (transfer_from): ", err)
		return err
	}

	if _, err := tx.Exec(ctx, "UPDATE users SET balance = balance + $1 WHERE id = $2", amount, toUserID); err != nil {
		r.logger.Errorln("Update user's <--  balance (transfer_to): ", err)
		return err
	}

	if _, err := tx.Exec(ctx, "INSERT INTO tx (user_id, amount, type) VALUES ($1, $2, $3)", fromUserID, amount, "transfer_out"); err != nil {
		r.logger.Errorln("Insert into tx (transfer_out): ", err)
		return err
	}

	if _, err := tx.Exec(ctx, "INSERT INTO tx (user_id, amount, type) VALUES ($1, $2, $3)", toUserID, amount, "transfer_in"); err != nil {
		r.logger.Errorln("Insert into tx (transfer_in): ", err)
		return err
	}

	errTx := tx.Commit(ctx)
	if errTx != nil {
		r.logger.Errorln("Commit tx's error: ", err)
		return err
	}
	r.logger.Infoln("TransferMoney: success")

	return nil
}

func (r *Repository) GetLastTxs(c context.Context, userID int) ([]entity.Tx, error) {
	ctx, cancel := context.WithCancel(c)
	defer cancel()

	rows, err := r.db.Query(ctx, "SELECT sender_name, receiver_name, amount, type, created_at FROM tx WHERE user_id = $1 ORDER BY created_at DESC LIMIT 10", userID)
	if err != nil {
		r.logger.Errorln("select (get last): ", err)
		return nil, err
	}
	defer rows.Close()

	var txs []entity.Tx
	for rows.Next() {
		var tx entity.Tx
		if err := rows.Scan(&tx.Sender_name, &tx.Receiver_name, &tx.Amount, &tx.Type, &tx.CreatedAt); err != nil {
			r.logger.Errorln(err)
			return nil, err
		}
		txs = append(txs, tx)
	}

	return txs, nil
}
