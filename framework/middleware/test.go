package middleware

import (
	"fmt"
	"github.com/lai416703504/jin/framework/gin"
)

func Test1() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("middleware pre test1")
		ctx.Next()
		fmt.Println("middleware post test1")
	}
}

func Test2() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("middleware pre test2")
		ctx.Next()
		fmt.Println("middleware post test2")
	}
}

func Test3() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("middleware pre test3")
		ctx.Next()
		fmt.Println("middleware post test3")
	}
}

func Test5() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("middleware pre test5")
		ctx.Next()
		fmt.Println("middleware post test5")
	}
}
