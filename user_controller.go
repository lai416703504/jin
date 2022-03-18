package main

import (
	"github.com/lai416703504/jin/framework"
	"time"
)

func UserLoginController(ctx *framework.Context) error {
	// 打印控制器名字
	ctx.SetOkStatus().Json("ok, UserLoginController")
	return nil
}

func UserWaitController(ctx *framework.Context) error {
	// 打印控制器名字
	time.Sleep(10 * time.Second)
	ctx.SetOkStatus().Json("ok, UserWait")
	return nil
}
