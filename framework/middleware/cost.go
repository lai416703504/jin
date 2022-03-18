package middleware

import (
	"github.com/lai416703504/jin/framework"
	"log"
	"strings"
	"time"
)

func Cost() framework.ControllerHandler {
	//使用回调函数
	return func(ctx *framework.Context) error {
		//记录开始时间
		start := time.Now()

		//使用next执行具体的业务逻辑
		ctx.Next()

		//记录结束时间
		end := time.Now()
		cost := end.Sub(start)
		log.Printf(
			"method: %v, api uri: %v, cost: %v s",
			strings.ToUpper(ctx.GetRequest().Method),
			ctx.GetRequest().RequestURI,
			cost.Seconds(),
		)

		return nil
	}
}
