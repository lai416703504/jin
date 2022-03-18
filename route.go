package main

import (
	"github.com/lai416703504/jin/framework/gin"
	"github.com/lai416703504/jin/framework/middleware"
)

func registerRouter(core *gin.Engine) {
	// core.Get("foo", framework.TimeoutHandler(FooControllerHandler, time.Second*1))
	//core.Get("/foo", FooControllerHandler)
	// 需求1+2:HTTP方法+静态路由匹配
	//core.Get("/user/login", UserLoginController)
	// 在group中使用middleware.Test3() 为单个路由增加中间件
	core.GET("/user/login", middleware.Test3(), UserLoginController)
	core.GET("/user/wait", UserWaitController)
	// 需求3:批量通用前缀
	subjectApi := core.Group("/subject")
	{
		subjectApi.Use(middleware.Test5())
		// 需求4:动态路由
		subjectApi.DELETE("/:id", SubjectDelController)
		subjectApi.PUT("/:id", SubjectUpdateController)
		// 在group中使用middleware.Test3() 为单个路由增加中间件
		subjectApi.GET("/:id", middleware.Test3(), SubjectGetController)
		subjectApi.GET("/list/all", SubjectListController)

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.GET("/name", SubjectNameController)
		}
	}

	//core.GetDfs()
}
