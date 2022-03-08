package mysql

import (
	"time"

	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	MysqlDB *gorm.DB
)

var (
	_defaultMaxOpenConnections    = 10
	_defaultMaxIdleConnections    = 10
	_defaultMaxConnectionLifeTime = time.Second * 10
)

type MySQL struct {
	maxOpenConnections    int
	maxIdleConnections    int
	maxConnectionLifeTime time.Duration
}

func InitMysqlClient(dsn string, opts ...Options) {
	var err error

	MysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction:                   false,
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   logger.Default.LogMode(logger.Warn),
		NamingStrategy:                           schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		panic(err)
	}

	db, err := MysqlDB.DB()
	if err != nil {
		panic(err)
	}

	MySQLClient := &MySQL{
		maxConnectionLifeTime: _defaultMaxConnectionLifeTime,
		maxOpenConnections:    _defaultMaxOpenConnections,
		maxIdleConnections:    _defaultMaxIdleConnections,
	}

	for _, opt := range opts {
		opt(MySQLClient)
	}

	db.SetMaxOpenConns(MySQLClient.maxOpenConnections)
	db.SetMaxIdleConns(MySQLClient.maxIdleConnections)
	db.SetConnMaxLifetime(MySQLClient.maxConnectionLifeTime)

}
