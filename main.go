package main

import (
	"github.com/lai416703504/jin/app/console"
	JinHttp "github.com/lai416703504/jin/app/http"
	"github.com/lai416703504/jin/framework"
	"github.com/lai416703504/jin/framework/provider/app"
	"github.com/lai416703504/jin/framework/provider/distributed"
	"github.com/lai416703504/jin/framework/provider/kernel"
)

func main() {
	// 初始化服务容器
	container := framework.NewJinContainer()

	// 绑定 App 服务提供者
	container.Bind(&app.JinAppProvider{})

	// 后续初始化需要绑定的服务提供者...
	container.Bind(&distributed.LocalDistributedProvider{})

	//将 HTTP 引擎初始化,并且作为服务提供者绑定到服务容器中
	if engine, err := JinHttp.NewHttpEngine(); err == nil {
		container.Bind(&kernel.JinKernelProvider{HttpEngine: engine})
	}

	console.RunCommand(container)

	//core := gin.New()
	//
	//// 绑定具体的服务
	//core.Bind(&app.JinAppProvider{})
	//core.Bind(&demo.DemoServiceProvider{})
	////core.Use(middleware.Recovery(), middleware.Cost(), middleware.Test1(), middleware.Test2())
	//core.Use(gin.Recovery())
	//core.Use(middleware.Cost())
	//JinHttp.Routes(core)
	//server := &http.Server{
	//	Handler: core,
	//	Addr:    ":8080",
	//}
	//
	//// 这个 Goroutine 是启动服务的 Goroutine
	//go func() {
	//	server.ListenAndServe()
	//}()
	//
	//// 当前的 Goroutine 等待信号量
	//quit := make(chan os.Signal)
	//// 监控信号：SIGINT, SIGTERM, SIGQUIT
	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	//// 这里会阻塞当前 Goroutine 等待信号
	//<-quit
	//
	//// 调用Server.Shutdown graceful结束
	//timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//
	//if err := server.Shutdown(timeoutCtx); err != nil {
	//	log.Fatal("Server Shutdown:", err)
	//}
}
