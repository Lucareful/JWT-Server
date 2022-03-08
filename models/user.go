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

func (u *User) Create() error {
	return ms.MysqlDB.Create(u).Error
}

func (u *User) Update() error {
	return ms.MysqlDB.Model(u).Updates(u).Error
}
