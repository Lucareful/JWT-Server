package models

import (
	"github.com/google/uuid"
	ms "github.com/luenci/oauth2/store/mysql"
	"gorm.io/gorm"
)

// User struct 用户信息表
type User struct {
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

func (u *User) GetUserByName(name string) error {
	return ms.MysqlDB.Where("name = ?", name).First(&u).Error
}

func (u *User) GetAllUsers() ([]User, error) {
	var users []User
	err := ms.MysqlDB.Select("name", "password").Find(&users).Error
	return users, err
}

func (u *User) GetUserID(name, password string) error {
	return ms.MysqlDB.Select("user_id").Where("name = ? and password = ?", name, password).Find(&u).Error
}

func (u *User) Create() error {
	return ms.MysqlDB.Create(&u).Error
}

func (u *User) Update() error {
	return ms.MysqlDB.Model(u).Updates(u).Error
}

func NewUser() *User {
	return &User{}
}
