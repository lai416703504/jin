package main

import (
	"github.com/lai416703504/jin/framework/gin"
	"time"
)

func UserLoginController(ctx *gin.Context) {
	// 打印控制器名字
	ctx.ISetOkStatus().IJson("ok, UserLoginController")

}

func UserWaitController(ctx *gin.Context) {
	// 打印控制器名字
	time.Sleep(10 * time.Second)
	ctx.ISetOkStatus().IJson("ok, UserWait")

}
