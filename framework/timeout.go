package framework

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func TimeoutHandler(fun ControllerHandler, d time.Duration) ControllerHandler {
	// 使用函数回调
	return func(ctx *Context) error {
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)

		durationCtx, cancel := context.WithTimeout(ctx.BaseContext(), 1*time.Second)
		defer cancel()

		ctx.request.WithContext(durationCtx)

		//mu:=sync.Mutex{}
		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()
			//Do real action
			fun(ctx)

			finish <- struct{}{}
		}()

		select {
		case p := <-panicChan:
			log.Println(p)
			ctx.responseWriter.WriteHeader(http.StatusInternalServerError)
		case <-finish:
			fmt.Println("finish")
		case <-durationCtx.Done():
			ctx.SetHasTimeout()
			ctx.responseWriter.Write([]byte("time out"))
		}

		return nil
	}
}
