package main

import (
	"context"
	"github.com/duanchuansong/stk/xlog"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"zj/app/config"
	"zj/app/server"
	"zj/app/service"
)

func main() {
	xlog.Init()
	cfg := config.New()
	s := service.New(cfg)

	httpServer := server.StartHttpServer(s)

	go server.StartStaticServer(cfg.Static)

	go func() {
		// 服务连接
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			xlog.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	xlog.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	xlog.Info("Server exiting")
}
