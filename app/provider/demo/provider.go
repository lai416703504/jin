package demo

import "github.com/lai416703504/jin/framework"

type DemoProvider struct {
	framework.ServiceProvider
	c framework.Container
}

func (sp *DemoProvider) Register(container framework.Container) framework.NewInstance {
	return NewService
}

func (sp *DemoProvider) Boot(container framework.Container) error {
	sp.c = container
	return nil
}

func (sp *DemoProvider) IsDefer() bool {
	return false
}

func (sp *DemoProvider) Params(container framework.Container) []interface{} {
	return []interface{}{sp.c}
}

func (sp *DemoProvider) Name() string {
	return DemoKey
}
