package repository

import (
	"github.com/jmoiron/sqlx"
)

const (
	accountTable = "account"
	taskTable    = "task"
)

// Config - db
type Config struct {
	Driver string
}

// NewSqlite3DB - open connect and ping trying
func NewSqlite3DB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open(cfg.Driver, ":memory:")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}