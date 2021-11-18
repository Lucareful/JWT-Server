package main

// 生成 swagger
//go:generate swag init -g routers/router.go --parseDependency --parseInternal --generatedTime --parseDepth 10
// 生成错误码
//go:generate codegen -type=int pkg/e/apiserver.go

//func main() {
//	config.InitConf()
//	conf := config.GetConf()
//
//	srv := &http.Server{
//		Addr:           conf.Server.BindAddress,
//		Handler:        routers.InitRouter(conf),
//		ReadTimeout:    conf.Server.ReadTimeout * time.Second,
//		WriteTimeout:   conf.Server.WriteTimeout * time.Second,
//		MaxHeaderBytes: 1 << 20,
//	}
//
//	go func() {
//		// service connections
//		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
//			log.Fatalf("listen: %s\n", err)
//		}
//	}()
//	store.PoolInitRedis(conf.Redis.Host, conf.Redis.Password,conf.Redis.DB)
//
//	log.Printf("server is runing: %s\n", conf.Server.BindAddress)
//
//	// Wait for interrupt signal to gracefully shut down the server with
//	// a timeout of 5 seconds.
//	quit := make(chan os.Signal, 1)
//	signal.Notify(quit,
//		syscall.SIGINT,
//		syscall.SIGTERM,
//		syscall.SIGSEGV,
//		syscall.SIGABRT,
//		syscall.SIGILL,
//		syscall.SIGFPE,
//		os.Interrupt)
//	<-quit
//	log.Println("Shutdown Server ...")
//
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//	if err := srv.Shutdown(ctx); err != nil {
//		log.Fatal("Server Shutdown:", err)
//	}
//	log.Println("Server exiting")
//	close(quit)
//}

import (
	"crypto/md5"
	"fmt"
)

func main() {
	data := []byte("123213213")
	fmt.Printf("%x", md5.Sum(data))
}
