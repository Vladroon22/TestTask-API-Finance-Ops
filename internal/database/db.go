package database

import (
	"context"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func DBConn() (*pgxpool.Pool, error) {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, os.Getenv("DB_URL"))
	if err != nil {
		return nil, err
	}
	if err := ping(ctx, pool); err != nil {
		return nil, err
	}
	return pool, nil
}

func ping(c context.Context, cn *pgxpool.Pool) error {
	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()
	var err error
	for i := 0; i < 5; i++ {
		if err = cn.Ping(ctx); err == nil {
			return nil
		}
		time.Sleep(time.Millisecond * 500)
	}
	return err
}
