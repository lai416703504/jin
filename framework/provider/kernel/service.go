package kernel

import (
	"github.com/lai416703504/jin/framework/gin"
	"net/http"
)

//引擎服务
type JinKernelService struct {
	engine *gin.Engine
}

// 初始化web引擎服务实例
func NewJinKernelService(params ...interface{}) (interface{}, error) {
	httpEngine := params[0].(*gin.Engine)
	return &JinKernelService{engine: httpEngine}, nil
}

// 返回web引擎
func (s *JinKernelService) HttpEngine() http.Handler {
	return s.engine
}
