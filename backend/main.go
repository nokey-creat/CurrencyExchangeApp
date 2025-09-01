package main

import (
	"CurrencyExchangeApp/config"
	"CurrencyExchangeApp/router"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	config.InitConfig()

	r := router.SetupRouter()

	port := config.AppConfig.App.Port
	if port == "" {
		port = ":8080" //默认配置
	}

	//优雅地启动和退出服务器

	srv := &http.Server{
		Addr:    port,
		Handler: r.Handler(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shut down server...")

	ctx, concel := context.WithTimeout(context.Background(), 5*time.Second)
	defer concel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("server shutdown: %s\n", err)
	}

	log.Println("Server exiting")
}
