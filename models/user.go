package models

import (
	ms "github.com/luenci/oauth2/store/mysql"
	"gorm.io/gorm"
)

// User struct 用户信息表
type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);not null"`
	Password string `gorm:"type:varchar(100);not null"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) GetUserByName(name string) error {
	return ms.MysqlDB.Where("name = ?", name).First(&u).Error
}

func (u *User) GetAllUsers() ([]User, error) {
	var users []User
	err := ms.MysqlDB.Select("name", "password").Find(&users).Error
	return users, err
}

func (u *User) Create() error {
	return ms.MysqlDB.Create(u).Error
}

func (u *User) Update() error {
	return ms.MysqlDB.Model(u).Updates(u).Error
}

func NewUser() *User {
	return &User{}
}
