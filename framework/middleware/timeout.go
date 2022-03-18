package middleware

import (
	"context"
	"fmt"
	"github.com/lai416703504/jin/framework/gin"
	"log"
	"net/http"
	"time"
)

func TimeoutHandler(d time.Duration) gin.HandlerFunc {
	// 使用函数回调
	return func(ctx *gin.Context) {
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)

		durationCtx, cancel := context.WithTimeout(ctx.BaseContext(), d)
		defer cancel()

		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()
			//Do real action
			ctx.Next()

			finish <- struct{}{}
		}()

		select {
		case p := <-panicChan:
			log.Println(p)
			ctx.ISetStatus(http.StatusInternalServerError).IJson("time out")
		case <-finish:
			fmt.Println("finish")
		case <-durationCtx.Done():
			ctx.ISetStatus(http.StatusInternalServerError).IJson("time out")
		}
	}
}
