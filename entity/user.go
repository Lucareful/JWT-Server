package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User struct 用户信息表
type User struct {
	// 绑定的数据源
	gorm.Model
	UserId   string `gorm:"type:varchar(36);unique_index;not null"`
	Name     string `gorm:"type:varchar(100);not null"`
	Password string `gorm:"type:varchar(100);not null"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Statement.SetColumn("user_id", uuid.New().String())
	return
}

func (u *User) TableName() string {
	return "user"
}

func NewUser() *User {
	return &User{}
}
