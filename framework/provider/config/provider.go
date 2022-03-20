package config

import (
	"github.com/lai416703504/jin/framework"
	"github.com/lai416703504/jin/framework/contract"
	"path/filepath"
)

type JinConfigProvider struct{}

func (j *JinConfigProvider) Register(container framework.Container) framework.NewInstance {
	return NewJinConfig
}

func (j *JinConfigProvider) Boot(container framework.Container) error {
	return nil
}

func (j *JinConfigProvider) IsDefer() bool {
	return false
}

func (j *JinConfigProvider) Params(container framework.Container) []interface{} {
	appService := container.MustMake(contract.AppKey).(contract.App)
	envService := container.MustMake(contract.EnvKey).(contract.Env)
	env := envService.AppEnv()
	//配置文件夹地址
	configFolder := appService.ConfigFolder()
	envFolder := filepath.Join(configFolder, env)
	return []interface{}{container, envFolder, envService.All()}
}

func (j *JinConfigProvider) Name() string {
	return contract.ConfigKey
}
