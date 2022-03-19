package app

import (
	"errors"
	"github.com/google/uuid"
	"github.com/lai416703504/jin/framework"
	"github.com/lai416703504/jin/framework/util"
	"path/filepath"
)

// JinApp 代表jin框架的App实现
type JinApp struct {
	container  framework.Container //服务容器
	baseFolder string              //基础路径
	appId      string              // 表示当前这个app的唯一id, 可以用于分布式锁等
}

func NewJinApp(param ...interface{}) (interface{}, error) {
	if len(param) != 2 {
		return nil, errors.New("param error")
	}

	// 有两个参数，一个是容器，一个是baseFolder
	container := param[0].(framework.Container)
	baseFolder := param[1].(string)
	appId := uuid.New().String()
	return &JinApp{container: container, baseFolder: baseFolder, appId: appId}, nil
}

func (j *JinApp) AppID() string {
	return j.appId
}

func (j *JinApp) Version() string {
	return "0.0.1"
}

func (j *JinApp) BaseFolder() string {

	if j.baseFolder != "" {
		return j.baseFolder
	}

	//// 如果没有设置，则使用参数
	//var baseFolder string
	//flag.StringVar(&baseFolder, "base_folder", "", "base_folder参数, 默认为当前路径")
	//flag.Parse()
	//if baseFolder != "" {
	//	return baseFolder
	//}

	// 如果参数也没有，使用默认的当前路径
	return util.GetExecDirectory()
}

func (j *JinApp) ConfigFolder() string {
	return filepath.Join(j.BaseFolder(), "config")
}

func (j *JinApp) StorageFolder() string {
	return filepath.Join(j.BaseFolder(), "storage")
}

func (j *JinApp) LogFolder() string {
	return filepath.Join(j.StorageFolder(), "log")
}

func (j *JinApp) HttpFolder() string {
	return filepath.Join(j.BaseFolder(), "http")
}

func (j *JinApp) ConsoleFolder() string {
	return filepath.Join(j.BaseFolder(), "console")
}

func (j *JinApp) ProviderFolder() string {
	return filepath.Join(j.BaseFolder(), "provider")
}

func (j *JinApp) MiddlewareFolder() string {
	return filepath.Join(j.HttpFolder(), "middleware")
}

func (j *JinApp) CommandFolder() string {
	return filepath.Join(j.ConsoleFolder(), "command")
}

func (j *JinApp) RuntimeFolder() string {
	return filepath.Join(j.StorageFolder(), "runtime")
}

func (j *JinApp) TestFolder() string {
	return filepath.Join(j.BaseFolder(), "test")
}
