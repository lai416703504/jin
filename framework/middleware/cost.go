package middleware

import (
	"github.com/lai416703504/jin/framework/gin"
	"log"
	"strings"
	"time"
)

func Cost() gin.HandlerFunc {
	//使用回调函数
	return func(ctx *gin.Context) {
		//记录开始时间
		start := time.Now()

		log.Printf("api uri start: %v", ctx.Request.RequestURI)

		//使用next执行具体的业务逻辑
		ctx.Next()

		//记录结束时间
		end := time.Now()
		cost := end.Sub(start)
		log.Printf(
			"method: %v, api uri: %v, cost: %v s",
			strings.ToUpper(ctx.Request.Method),
			ctx.Request.RequestURI,
			cost.Seconds(),
		)
	}
}
