package main

import "github.com/lai416703504/jin/framework"

func SubjectDelController(ctx *framework.Context) error {
	// 打印控制器名字
	ctx.SetOkStatus().Json("ok, SubjectDelController")
	return nil
}

func SubjectUpdateController(ctx *framework.Context) error {
	// 打印控制器名字
	ctx.SetOkStatus().Json("ok, SubjectUpdateController")
	return nil
}

func SubjectGetController(ctx *framework.Context) error {
	// 打印控制器名字
	ctx.SetOkStatus().Json("ok, SubjectGetController")
	return nil
}

func SubjectListController(ctx *framework.Context) error {
	// 打印控制器名字
	ctx.SetOkStatus().Json("ok, SubjectListController")
	return nil
}

func SubjectNameController(ctx *framework.Context) error {
	// 打印控制器名字
	ctx.SetOkStatus().Json("ok, SubjectNameController")
	return nil
}
