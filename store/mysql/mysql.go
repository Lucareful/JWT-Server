package mysql

import (
	"sync"

	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	once  sync.Once
	dbIns *gorm.DB
)

func NewMysqlStore(dsn string, opts ...Options) (*gorm.DB, error) {

	var err error
	once.Do(func() {
		dbIns, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			SkipDefaultTransaction:                   false,
			DisableForeignKeyConstraintWhenMigrating: true,
			Logger:                                   logger.Default.LogMode(logger.Warn),
			NamingStrategy:                           schema.NamingStrategy{SingularTable: true},
		})
		if err != nil {
			return
		}
		db, _ := dbIns.DB()
		for _, opt := range opts {
			opt(db)
		}
		return
	})

	return dbIns, err
}
