package kernel

import (
	"github.com/lai416703504/jin/framework"
	"github.com/lai416703504/jin/framework/contract"
	"github.com/lai416703504/jin/framework/gin"
)

// JinKernelProvider 提供web引擎
type JinKernelProvider struct {
	HttpEngine *gin.Engine
}

func (j *JinKernelProvider) Register(container framework.Container) framework.NewInstance {
	return NewJinKernelService
}

func (j *JinKernelProvider) Boot(container framework.Container) error {
	if j.HttpEngine == nil {
		j.HttpEngine = gin.Default()
	}
	j.HttpEngine.SetContainer(container)

	return nil
}

func (j *JinKernelProvider) IsDefer() bool {
	return false
}

func (j *JinKernelProvider) Params(container framework.Container) []interface{} {
	return []interface{}{j.HttpEngine}
}

func (j *JinKernelProvider) Name() string {
	return contract.KernelKey
}
