package middleware

import (
	"github.com/lai416703504/jin/framework"
	"net/http"
)

func Recovery() framework.ControllerHandler {
	//使用回调函数
	return func(ctx *framework.Context) error {
		// 核心在增加这个recover机制，捕获c.Next()出现的panic
		defer func() {
			if err := recover(); err != nil {
				ctx.SetStatus(http.StatusInternalServerError).Json(err)
			}
		}()

		ctx.Next()

		return nil
	}
}
