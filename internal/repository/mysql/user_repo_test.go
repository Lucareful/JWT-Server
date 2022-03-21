package mysql

import (
	"context"
	"fmt"
	"testing"

	"github.com/luenci/oauth2/config"

	"github.com/luenci/oauth2/store/mysql"

	"github.com/luenci/oauth2/entity"
)

func TestUserRepository_Create(t *testing.T) {
	type args struct {
		ctx  context.Context
		user entity.User
	}
	tests := []struct {
		name    string
		args    args
		want    entity.User
		wantErr bool
	}{
		{"Create", args{ctx: context.Background(), user: entity.User{
			Name:     "luenci",
			Password: "admin",
		}}, entity.User{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config.InitConf()
			conf := config.GetConf()
			dbins, _ := mysql.NewMysqlStore(conf.Mysql.DSN)
			u := &UserRepository{
				db: dbins,
			}
			got, _ := u.Create(tt.args.ctx, tt.args.user)
			fmt.Printf("%+v\n", got)
		})
	}
}

func TestUserRepository_GetUserID(t *testing.T) {
	type args struct {
		ctx      context.Context
		name     string
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    entity.User
		wantErr bool
	}{
		{"Create", args{ctx: context.Background(), name: "luenci", password: "admin"}, entity.User{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config.InitConf()
			conf := config.GetConf()
			dbins, _ := mysql.NewMysqlStore(conf.Mysql.DSN)
			u := &UserRepository{
				db: dbins,
			}
			got, _ := u.GetUserID(tt.args.ctx, tt.args.name, tt.args.password)
			fmt.Printf("%+v\n", got)
		})
	}
}

func TestUserRepository_GetAllUsers(t *testing.T) {

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    []entity.User
		wantErr bool
	}{
		{"GetAllUsers", args{ctx: context.Background()}, []entity.User{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			config.InitConf()
			conf := config.GetConf()
			dbins, _ := mysql.NewMysqlStore(conf.Mysql.DSN)
			u := &UserRepository{
				db: dbins,
			}
			got, _ := u.GetAllUsers(tt.args.ctx)
			fmt.Printf("%+v\n", got)
		})
	}
}
