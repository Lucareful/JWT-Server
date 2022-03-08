package mysql

import "time"

type Options func(m *MySQL)

func WithMaxOpenConnections(maxOpenConnections int) Options {
	return func(m *MySQL) {
		m.maxOpenConnections = maxOpenConnections
	}
}

func WithMaxIdleConnections(maxIdleConnections int) Options {
	return func(m *MySQL) {
		m.maxIdleConnections = maxIdleConnections
	}
}

func WithMaxConnectionLifeTime(maxConnectionLifeTime time.Duration) Options {
	return func(m *MySQL) {
		m.maxConnectionLifeTime = maxConnectionLifeTime
	}
}
