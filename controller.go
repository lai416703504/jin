package main

import (
	"context"
	"fmt"
	"github.com/lai416703504/jin/framework"
	"log"
	"net/http"
	"time"
)

func FooControllerHandler(ctx *framework.Context) error {
	//return ctx.Json(200, map[string]interface{}{"code": 0})
	finish := make(chan struct{}, 1)
	panicChan := make(chan interface{}, 1)

	durationCtx, cancel := context.WithTimeout(ctx.BaseContext(), 1*time.Second)
	defer cancel()

	//mu:=sync.Mutex{}
	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		//Do real action
		time.Sleep(10 * time.Second)
		ctx.SetOkStatus().Json(http.StatusText(http.StatusOK))
		finish <- struct{}{}
	}()

	select {
	case p := <-panicChan:
		ctx.WriterMux().Lock()
		defer ctx.WriterMux().Unlock()
		log.Println(p)
		ctx.SetStatus(http.StatusInternalServerError).Json("panic")
	case <-finish:
		fmt.Println("finish")
	case <-durationCtx.Done():
		ctx.WriterMux().Lock()
		defer ctx.WriterMux().Unlock()
		ctx.SetStatus(http.StatusInternalServerError).Json("time out")
		ctx.SetHasTimeout()
	}

	return nil
}
