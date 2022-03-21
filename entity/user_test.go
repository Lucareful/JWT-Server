package entity

import (
	"testing"

	"github.com/luenci/oauth2/config"

	"github.com/luenci/oauth2/store/mysql"
)

func TestUser_Create(t *testing.T) {
	type fields struct {
		Name     string
		Password string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{name: "create_user", fields: fields{Name: "test", Password: "test"}},
	}
	config.InitConf()
	conf := config.GetConf()

	mysql.InitMysqlClient(conf.Mysql.DSN, mysql.WithMaxIdleConnections(conf.Mysql.MaxIdleConnections),
		mysql.WithMaxOpenConnections(conf.Mysql.MaxOpenConnections),
		mysql.WithMaxConnectionLifeTime(conf.Mysql.MaxConnectionLifeTime))

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usr := NewUser()
			usr.Name = tt.fields.Name
			usr.Password = tt.fields.Password
			if err := usr.Create(); err != nil {
				t.Errorf("Create() error = %v", err)
			}
		})
	}
}
