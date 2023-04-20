package db

import "github.com/jackc/pgx/v5"

type DB struct {
	Conn *pgx.Conn
}

func NewDB() (*DB, error) {
	conn, err := pgx.Connect()
}
