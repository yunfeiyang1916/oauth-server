package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/yunfeiyang1916/toolkit/framework"

	"github.com/yunfeiyang1916/oauth-server/conf"
	httpserver "github.com/yunfeiyang1916/toolkit/framework/http/server"

	"github.com/yunfeiyang1916/oauth-server/service"

	"github.com/yunfeiyang1916/toolkit/logging"
)

var (
	server     http.Server
	httpServer httpserver.Server
)

// Init 启动http服务
func Init(srv *service.Container) {
	handler := newHttpHandler(srv)
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", conf.GlobalConfig.ServerConfig.Port),
		Handler: handler,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			logging.Fatalf("http server start failed, err %v", err)
		}
	}()
}

func InitFrm(s *service.Container) {
	srv = s
	httpServer = framework.HTTPServer()
	// add namespace plugin
	// httpServer.Use(httpplugin.Namespace)

	// register handler with http route
	initRouter(httpServer)
	// start a http server
	go func() {
		if err := httpServer.Run(); err != nil {
			logging.Fatalf("http server start failed, err %v", err)
		}
	}()
}

// Shutdown 关闭
func Shutdown(ctx context.Context) error {
	return server.Shutdown(ctx)
}
