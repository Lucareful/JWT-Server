package main

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	ms "github.com/luenci/oauth2/store/mysql"

	validate "github.com/luenci/oauth2/pkg/validator"

	"github.com/luenci/oauth2/store/redis"

	"github.com/luenci/oauth2/config"
	"github.com/luenci/oauth2/routers"
)

// 生成 swagger
//go:generate swag init -g routers/router.go --parseDependency --parseInternal --generatedTime --parseDepth 10
// 生成错误码
//go:generate codegen -type=int pkg/e/apiserver.go

func main() {
	// 随机数种子，用于生成随机数。
	rand.Seed(time.Now().UnixNano())

	// 初始化配置
	config.InitConf()
	conf := config.GetConf()

	// check 创建数据库连接
	if _, err := ms.NewMysqlStore(conf.Mysql.DSN,
		ms.WithMaxIdleConnections(conf.Mysql.MaxIdleConnections),
		ms.WithMaxOpenConnections(conf.Mysql.MaxOpenConnections),
		ms.WithMaxConnectionLifeTime(conf.Mysql.MaxConnectionLifeTime)); err != nil {
		panic(err)
	}

	// 创建redis连接
	redis.PoolInitRedis(conf.Redis.Host, conf.Redis.Password)

	srv := &http.Server{
		Addr:         conf.Server.BindAddress,
		Handler:      routers.InitRouter(conf),
		ReadTimeout:  conf.Server.ReadTimeout * time.Second,
		WriteTimeout: conf.Server.WriteTimeout * time.Second,
	}

	// 初始化翻译
	validate.Init()

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Printf("server is runing: %s\n", conf.Server.BindAddress)

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGSEGV,
		syscall.SIGABRT,
		syscall.SIGILL,
		syscall.SIGFPE,
		os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
	close(quit)
}
