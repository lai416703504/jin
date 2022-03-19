package http

import (
	"github.com/lai416703504/jin/app/http/module/demo"
	"github.com/lai416703504/jin/framework/gin"
)

func Routes(r *gin.Engine) {
	r.Static("/dist/", "./dist/")
	demo.Register(r)
}
