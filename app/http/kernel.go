package http

import "github.com/lai416703504/jin/framework/gin"

// NewHttpEngine is command
func NewHttpEngine() (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	Route(r)

	return r, nil
}
