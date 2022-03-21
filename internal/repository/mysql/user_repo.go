package mysql

import (
	"context"

	"github.com/luenci/oauth2/config"
	"github.com/luenci/oauth2/store/mysql"

	"github.com/luenci/oauth2/internal/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (u *UserRepository) GetUserByName(ctx context.Context, name string) (entity.User, error) {
	var user entity.User
	return user, u.db.Where("name = ?", name).First(&user).Error
}

func (u *UserRepository) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	return users, u.db.Select("name", "password").Find(&users).Error
}

func (u *UserRepository) GetUserID(ctx context.Context, name, password string) (entity.User, error) {
	var user entity.User
	return user, u.db.Select("user_id").Where("name = ? and password = ?", name, password).Find(&user).Error
}

func (u *UserRepository) Create(ctx context.Context, user entity.User) (entity.User, error) {
	return user, u.db.Create(&user).Error
}

func (u *UserRepository) Update(ctx context.Context) (entity.User, error) {
	var user entity.User
	return user, u.db.Model(user).Updates(user).Error
}

func NewUserRepository() *UserRepository {
	config.InitConf()
	conf := config.GetConf()
	dbIns, _ := mysql.NewMysqlStore(conf.Mysql.DSN)
	return &UserRepository{dbIns}
}
