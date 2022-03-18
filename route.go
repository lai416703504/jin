package main

import (
	"github.com/lai416703504/jin/framework"
	"github.com/lai416703504/jin/framework/middleware"
	"time"
)

func registerRouter(core *framework.Core) {
	// core.Get("foo", framework.TimeoutHandler(FooControllerHandler, time.Second*1))
	//core.Get("/foo", FooControllerHandler)
	// 需求1+2:HTTP方法+静态路由匹配
	//core.Get("/user/login", UserLoginController)
	// 在group中使用middleware.Test3() 为单个路由增加中间件
	core.Get("/user/login", middleware.Test3(), framework.TimeoutHandler(UserLoginController, time.Second))
	core.Get("/user/wait", UserWaitController)
	// 需求3:批量通用前缀
	subjectApi := core.Group("/subject")
	{
		subjectApi.Use(middleware.Test5())
		// 需求4:动态路由
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		// 在group中使用middleware.Test3() 为单个路由增加中间件
		subjectApi.Get("/:id", middleware.Test3(), SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.Get("/name", SubjectNameController)
		}
	}

	//core.GetDfs()
}
