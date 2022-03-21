package mysql

import (
	"sync"

	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	once sync.Once
)

func NewMysqlStore(dsn string, opts ...Options) (*gorm.DB, error) {

	var err error
	var dbIns *gorm.DB
	once.Do(func() {
		dbIns, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			SkipDefaultTransaction:                   false,
			DisableForeignKeyConstraintWhenMigrating: true,
			Logger:                                   logger.Default.LogMode(logger.Warn),
			NamingStrategy:                           schema.NamingStrategy{SingularTable: true},
		})
		if err != nil {
			panic(err)
		}
		db, _ := dbIns.DB()
		for _, opt := range opts {
			opt(db)
		}

		// uncomment the following line if you need auto migration the given repository
		// not suggested in production environment.
		// migrateDatabase(dbIns)

	})

	return dbIns, nil
}
