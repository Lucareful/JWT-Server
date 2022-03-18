package mysql

import (
	"sync"

	"github.com/luenci/oauth2/service"

	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	once sync.Once
)

type datastore struct {

	// can include two database instance if needed
	db *gorm.DB
}

func (ds *datastore) JWT() service.JWTService {
	return service.NewJWTServices(ds)
}

func (ds *datastore) Authorization() service.AuthorizationService {
	return (ds)
}

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

		// uncomment the following line if you need auto migration the given models
		// not suggested in production environment.
		// migrateDatabase(dbIns)

		mysqlFactory = &datastore{dbIns}
	})

	return mysqlFactory, nil
}
