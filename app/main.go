package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yunfeiyang1916/oauth-server/service"
	"github.com/yunfeiyang1916/oauth-server/transport/http"
)

func main() {
	srv := service.New()
	defer srv.Close()
	http.Init(srv)

	sigChan := make(chan os.Signal, 1)
	// 监听退出信号
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	s := <-sigChan
	log.Printf("get a signal %s\n", s.String())
	switch s {
	case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
		log.Println("server exit now...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := http.Shutdown(ctx); err != nil {
			log.Println("server shutdown failed")
		}
		log.Println("server exited")
	case syscall.SIGHUP:
	}
}
