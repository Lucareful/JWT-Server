package oauth2

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/luenci/oauth2/config"
	"github.com/luenci/oauth2/routers"
)

// 生成 swagger
//go:generate swag init -g routers/router.go --parseDependency --parseInternal --generatedTime --parseDepth 10
// 生成错误码
//go:generate codegen -type=int pkg/e/apiserver.go

func main() {
	config.InitConf()
	conf := config.GetConf()

	router := routers.InitRouter(conf)

	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", conf.Server.BindAddress),
		Handler:        router,
		ReadTimeout:    conf.Server.ReadTimeout,
		WriteTimeout:   conf.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		// service connections
		log.Printf("server is runing: %d\n", conf.Server.BindAddress)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}

	}()

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

}
