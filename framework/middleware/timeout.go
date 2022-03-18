package middleware

import (
	"context"
	"fmt"
	"github.com/lai416703504/jin/framework"
	"log"
	"net/http"
	"time"
)

func TimeoutHandler(d time.Duration) framework.ControllerHandler {
	// 使用函数回调
	return func(ctx *framework.Context) error {
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
			ctx.SetStatus(http.StatusInternalServerError).Json("time out")
		case <-finish:
			fmt.Println("finish")
		case <-durationCtx.Done():
			ctx.SetHasTimeout()
			ctx.SetStatus(http.StatusInternalServerError).Json("time out")
		}

		return nil
	}
}
