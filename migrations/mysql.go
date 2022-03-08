package main

import (
	"fmt"

	"github.com/luenci/oauth2/config"
	"github.com/luenci/oauth2/models"
	"github.com/luenci/oauth2/store/mysql"
)

func main() {
	// 初始化配置
	config.InitConf()
	conf := config.GetConf()

	fmt.Println(conf.Mysql.DSN)
	mysql.InitMysqlClient(conf.Mysql.DSN)
	err := mysql.MysqlDB.AutoMigrate(&models.User{}, &models.Client{}, &models.Token{})
	if err != nil {
		fmt.Println("AutoMigrate error:", err)
	}

}
