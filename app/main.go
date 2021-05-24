package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yunfeiyang1916/toolkit/framework"

	"github.com/yunfeiyang1916/toolkit/logging"

	"github.com/yunfeiyang1916/oauth-server/conf"

	"github.com/yunfeiyang1916/oauth-server/service"
	"github.com/yunfeiyang1916/oauth-server/transport/http"
)

func init() {
	configS := flag.String("config", "config/config.toml", "Configuration file")
	appS := flag.String("app", "", "App dir")
	flag.Parse()

	framework.Init(framework.ConfigPath(*configS))

	if *appS != "" {
		framework.InitNamespace(*appS)
	}

}

func main() {
	defer framework.Shutdown()

	// 初始化配置
	_, err := conf.Init()
	if err != nil {
		logging.Fatalf("service config init error %s", err)
	}

	srv := service.New()
	defer srv.Close()
	//http.Init(srv)
	http.InitFrm(srv)
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
