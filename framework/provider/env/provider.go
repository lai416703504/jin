package env

import (
	"github.com/lai416703504/jin/framework"
	"github.com/lai416703504/jin/framework/contract"
)

type JinEnvProvider struct {
	Folder string
}

func (j *JinEnvProvider) Register(container framework.Container) framework.NewInstance {
	return NewJinEnv
}

func (j *JinEnvProvider) Boot(container framework.Container) error {
	app := container.MustMake(contract.AppKey).(contract.App)
	j.Folder = app.BaseFolder()
	return nil
}

func (j *JinEnvProvider) IsDefer() bool {
	return false
}

func (j *JinEnvProvider) Params(container framework.Container) []interface{} {
	return []interface{}{j.Folder}
}

func (j *JinEnvProvider) Name() string {
	return contract.EnvKey
}
