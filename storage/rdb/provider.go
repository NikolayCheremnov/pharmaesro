package rdb

import (
	"database/sql"

	// TODO: fix linter message
	//nolint:blank-imports
	_ "github.com/lib/pq"
)

type DB struct {
	Session *sql.DB
}

func New(driverName string, connectionString string) (*DB, error) {
	db, err := sql.Open(driverName, connectionString)
	if err != nil {
		return nil, err
	}
	// TODO: add backoff
	if er := db.Ping(); er != nil {
		return nil, er
	}
	return &DB{Session: db}, nil
}

func (db *DB) Close() error {
	return db.Session.Close()
}
