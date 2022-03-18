package mysql

import (
	"database/sql"
	"time"
)

var (
	_defaultMaxOpenConnections    = 10
	_defaultMaxIdleConnections    = 10
	_defaultMaxConnectionLifeTime = time.Second * 10
)

type Options func(db *sql.DB)

func WithMaxOpenConnections(maxOpenConnections int) Options {
	return func(db *sql.DB) {
		if maxOpenConnections > 0 {
			db.SetMaxOpenConns(maxOpenConnections)
		}
		db.SetMaxOpenConns(_defaultMaxOpenConnections)
	}
}

func WithMaxIdleConnections(maxIdleConnections int) Options {
	return func(db *sql.DB) {
		if maxIdleConnections > 0 {
			db.SetMaxIdleConns(maxIdleConnections)
		}
		db.SetMaxIdleConns(_defaultMaxIdleConnections)
	}
}

func WithMaxConnectionLifeTime(maxConnectionLifeTime time.Duration) Options {
	return func(db *sql.DB) {
		if maxConnectionLifeTime > 0 {
			db.SetConnMaxLifetime(maxConnectionLifeTime)
		}
		db.SetConnMaxLifetime(_defaultMaxConnectionLifeTime)
	}
}
