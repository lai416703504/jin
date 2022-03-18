package app

import (
	"github.com/lai416703504/jin/framework"
	"github.com/lai416703504/jin/framework/contract"
)

type JinAppProvider struct {
	BaseFolder string
}

func (j *JinAppProvider) Register(container framework.Container) framework.NewInstance {
	return NewJinApp
}

func (j *JinAppProvider) Boot(container framework.Container) error {
	return nil
}

func (j *JinAppProvider) IsDefer() bool {
	return false
}

func (j *JinAppProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container, j.BaseFolder}
}

func (j *JinAppProvider) Name() string {
	return contract.AppKey
}
