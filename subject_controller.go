package main

import (
	"github.com/lai416703504/jin/framework/gin"
	"github.com/lai416703504/jin/provider/demo"
)

func SubjectDelController(ctx *gin.Context) {
	// 打印控制器名字
	ctx.ISetOkStatus().IJson("ok, SubjectDelController")

}

func SubjectUpdateController(ctx *gin.Context) {
	// 打印控制器名字
	ctx.ISetOkStatus().IJson("ok, SubjectUpdateController")

}

func SubjectGetController(ctx *gin.Context) {
	// 打印控制器名字
	ctx.ISetOkStatus().IJson("ok, SubjectGetController")

}

func SubjectListController(ctx *gin.Context) {

	// 获取 demo 服务实例
	demoService := ctx.MustMake(demo.Key).(demo.Service)

	//调用服务实例的方法
	foo := demoService.GetFoo()

	//输出结果
	ctx.ISetOkStatus().IJson(foo)

	//ctx.ISetOkStatus().IJson("ok, SubjectListController")

}

func SubjectNameController(ctx *gin.Context) {
	// 打印控制器名字
	ctx.ISetOkStatus().IJson("ok, SubjectNameController")

}
