package database

import (
	"github.com/jackc/pgx"
)

func DBConnection(dbUrl string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(pgx.ConnConfig{Host: dbUrl})
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func CloseConn(conn *pgx.Conn) error {
	return conn.Close()
}
