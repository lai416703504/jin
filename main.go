package main

import (
	"context"
	JinHttp "github.com/lai416703504/jin/app/http"
	"github.com/lai416703504/jin/framework/gin"
	"github.com/lai416703504/jin/framework/middleware"
	"github.com/lai416703504/jin/framework/provider/app"
	"github.com/lai416703504/jin/provider/demo"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	core := gin.New()

	// 绑定具体的服务
	core.Bind(&app.JinAppProvider{})
	core.Bind(&demo.DemoServiceProvider{})
	//core.Use(middleware.Recovery(), middleware.Cost(), middleware.Test1(), middleware.Test2())
	core.Use(gin.Recovery())
	core.Use(middleware.Cost())
	JinHttp.Route(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8080",
	}

	// 这个 Goroutine 是启动服务的 Goroutine
	go func() {
		server.ListenAndServe()
	}()

	// 当前的 Goroutine 等待信号量
	quit := make(chan os.Signal)
	// 监控信号：SIGINT, SIGTERM, SIGQUIT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// 这里会阻塞当前 Goroutine 等待信号
	<-quit

	// 调用Server.Shutdown graceful结束
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
