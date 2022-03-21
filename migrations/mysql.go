package main

import (
	"fmt"

	"github.com/luenci/oauth2/config"
	"github.com/luenci/oauth2/repository"
	"github.com/luenci/oauth2/store/mysql"

	"github.com/cheggaaa/pb/v3"
)

func main() {
	// 初始化配置
	config.InitConf()
	conf := config.GetConf()
	mysql.InitMysqlClient(conf.Mysql.DSN)
	fmt.Println("开始执行 migrations")
	bar := pb.StartNew(1)
	err := mysql.MysqlDB.AutoMigrate(&repository.User{}, &repository.Client{}, &repository.Token{})
	if err != nil {
		fmt.Println("AutoMigrate error:", err)
	}
	bar.Increment()

	bar.Finish()
	fmt.Println("执行 migrations 完毕")
}
