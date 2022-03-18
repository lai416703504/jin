package framework

import (
	"errors"
	"sync"
)

// Container 是一个服务容器，提供绑定服务和获取服务的功能
type Container interface {
	// Bind 绑定一个服务提供者，如果关键字凭证已经存在，会进行替换操作，不返回 error
	Bind(provider ServiceProvider) error
	// IsBind 关键字凭证是否已经绑定服务提供者
	IsBind(key string) bool

	// Make 根据关键字凭证获取一个服务，
	Make(key string) (interface{}, error)
	// MustMake 根据关键字凭证获取一个服务，如果这个关键字凭证未绑定服务提供者，那么会 panic。
	//所以在使用这个接口的时候请保证服务容器已经为这个关键字凭证绑定了服务提供者。
	MustMake(key string) interface{}
	//MakeNew 根据关键字凭证获取一个服务，只是这个服务并不是单例模式的
	//它是根据服务提供者注册的启动函数和传递的 params 参数实例化出来的
	//这个函数在需要为不同参数启动不同实例的时候非常有用
	MakeNew(key string, params []interface{}) (interface{}, error)
}

// JinContainer 是服务容器的具体实现
type JinContainer struct {
	Container // 强制要求 JinContainer 实现 Container 接口

	// providers 存储注册的服务提供者，key 为字符串凭证
	providers map[string]ServiceProvider

	// instances 存储具体的实例，key 为字符串凭证
	instances map[string]interface{}

	// lock 用于锁住对容器的变更操作 读多写少，用读写锁就可以了
	lock sync.RWMutex
}

// NewJinContainer 创建一个服务容器
func NewJinContainer() *JinContainer {
	return &JinContainer{
		providers: map[string]ServiceProvider{},
		instances: map[string]interface{}{},
		lock:      sync.RWMutex{},
	}
}

func (j *JinContainer) Bind(provider ServiceProvider) error {
	j.lock.Lock()
	defer j.lock.Unlock()

	key := provider.Name()

	j.providers[key] = provider

	// if provider is not defer
	if provider.IsDefer() == false {
		if err := provider.Boot(j); err != nil {
			return err
		}
		//实例化方法
		params := provider.Params(j)
		method := provider.Register(j)
		instance, err := method(params...)
		if err != nil {
			return err
		}
		j.instances[key] = instance
	}

	return nil
}

func (j *JinContainer) IsBind(key string) bool {
	return j.findServiceProvider(key) != nil
}

func (j *JinContainer) findServiceProvider(key string) ServiceProvider {
	j.lock.RLock()
	defer j.lock.RUnlock()
	if sp, ok := j.providers[key]; ok {
		return sp
	}
	return nil
}

func (j *JinContainer) Make(key string) (interface{}, error) {
	return j.make(key, nil, false)
}

func (j *JinContainer) MustMake(key string) interface{} {

	serv, err := j.make(key, nil, false)

	if err != nil {
		panic(err)
	}

	return serv
}

func (j *JinContainer) MakeNew(key string, params []interface{}) (interface{}, error) {
	return j.make(key, params, true)
}

//真正的实例化一个服务
func (j *JinContainer) make(key string, params []interface{}, froceNew bool) (interface{}, error) {
	j.lock.RLock()
	defer j.lock.RUnlock()
	//查询是否已经注册了这个服务提供者，如果没有注册，则返回错误
	sp := j.findServiceProvider(key)
	if sp == nil {
		return nil, errors.New("contract " + key + " have not register")
	}

	if froceNew {
		return j.newInstance(sp, params)
	}

	// 不需要强制重新实例化，如果容器中已经实例化了，那么就直接使用容器中的实例
	if ins, ok := j.instances[key]; ok {
		return ins, nil
	}

	// 容器中还未实例化，则进行一次实例化
	inst, err := j.newInstance(sp, params)
	if err != nil {
		return nil, err
	}

	j.instances[key] = inst
	return inst, nil

}

func (j *JinContainer) newInstance(sp ServiceProvider, params []interface{}) (interface{}, error) {
	// force new a
	if err := sp.Boot(j); err != nil {
		return nil, err
	}
	if params == nil {
		params = sp.Params(j)
	}
	method := sp.Register(j)
	ins, err := method(params...)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return ins, err
}
