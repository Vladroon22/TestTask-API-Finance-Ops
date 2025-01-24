package repository

import (
	"context"

	"github.com/Vladroon22/TestTask-Bank-Operation/internal/entity"
	"github.com/jackc/pgx"
)

type InterfaceRepo interface {
	TopUpBalance(userID int, amount float64) error
	TransferMoney(fromUserID, toUserID int, amount float64) error
	GetLastTransactions(userID int) ([]entity.Transaction, error)
}

type Repository struct {
	db *pgx.Conn
}

func NewFinanceRepository(db *pgx.Conn) *Repository {
	return &Repository{db: db}
}

func (r *Repository) TopUpBalance(userID int, amount float64) error {
	ctx := context.Background()
	tx, err := r.db.BeginEx(ctx, &pgx.TxOptions{IsoLevel: pgx.ReadUncommitted})
	if err != nil {
		return err
	}
	defer tx.RollbackEx(ctx)

	_, err = tx.ExecEx(ctx, "UPDATE users SET balance = balance + $1 WHERE id = $2", &pgx.QueryExOptions{}, amount, userID)
	if err != nil {
		return err
	}

	_, err = tx.ExecEx(ctx, "INSERT INTO transactions (user_id, amount, type) VALUES ($1, $2, 'topup')", &pgx.QueryExOptions{}, userID, amount)
	if err != nil {
		return err
	}

	return tx.CommitEx(ctx)
}

func (r *Repository) TransferMoney(fromUserID, toUserID int, amount float64) error {
	ctx := context.Background()
	tx, err := r.db.BeginEx(ctx, &pgx.TxOptions{IsoLevel: pgx.ReadUncommitted})
	if err != nil {
		return err
	}
	defer tx.RollbackEx(ctx)

	_, err = tx.ExecEx(ctx, "UPDATE users SET balance = balance - $1 WHERE id = $2", &pgx.QueryExOptions{}, amount, fromUserID)
	if err != nil {
		return err
	}

	_, err = tx.ExecEx(ctx, "UPDATE users SET balance = balance + $1 WHERE id = $2", &pgx.QueryExOptions{}, amount, toUserID)
	if err != nil {
		return err
	}

	_, err = tx.ExecEx(ctx, "INSERT INTO transactions (user_id, amount, type) VALUES ($1, $2, 'transfer_out')", &pgx.QueryExOptions{}, fromUserID, amount)
	if err != nil {
		return err
	}

	_, err = tx.ExecEx(ctx, "INSERT INTO transactions (user_id, amount, type) VALUES ($1, $2, 'transfer_in')", &pgx.QueryExOptions{}, toUserID, amount)
	if err != nil {
		return err
	}

	return tx.CommitEx(ctx)
}

func (r *Repository) GetLastTransactions(userID int) ([]entity.Transaction, error) {
	ctx := context.Background()
	rows, err := r.db.QueryEx(ctx, "SELECT id, user_id, amount, type, created_at FROM transactions WHERE user_id = $1 ORDER BY created_at DESC LIMIT 10", &pgx.QueryExOptions{}, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []entity.Transaction
	for rows.Next() {
		var t entity.Transaction
		if err := rows.Scan(&t.ID, &t.UserID, &t.Amount, &t.Type, &t.CreatedAt); err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}

	return transactions, nil
}
