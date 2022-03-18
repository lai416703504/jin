package app

import (
	"errors"
	"flag"
	"github.com/lai416703504/jin/framework"
	"github.com/lai416703504/jin/framework/util"
	"path/filepath"
)

// JinApp 代表jin框架的App实现
type JinApp struct {
	container  framework.Container //服务容器
	baseFolder string              //基础路径
}

func NewJinApp(param ...interface{}) (interface{}, error) {
	if len(param) != 2 {
		return nil, errors.New("param error")
	}

	// 有两个参数，一个是容器，一个是baseFolder
	container := param[0].(framework.Container)
	baseFolder := param[1].(string)

	return &JinApp{container: container, baseFolder: baseFolder}, nil
}

func (j *JinApp) Version() string {
	return "0.0.1"
}

func (j *JinApp) BaseFloder() string {

	if j.baseFolder != "" {
		return j.baseFolder
	}

	// 如果没有设置，则使用参数
	var baseFolder string
	flag.StringVar(&baseFolder, "base_folder", "", "base_folder参数, 默认为当前路径")
	flag.Parse()
	if baseFolder != "" {
		return baseFolder
	}

	// 如果参数也没有，使用默认的当前路径
	return util.GetExecDirectory()
}

func (j *JinApp) ConfigFloder() string {
	return filepath.Join(j.BaseFloder(), "config")
}

func (j *JinApp) StorageFolder() string {
	return filepath.Join(j.BaseFloder(), "storage")
}

func (j *JinApp) LogFloder() string {
	return filepath.Join(j.StorageFolder(), "log")
}

func (j *JinApp) HttpFolder() string {
	return filepath.Join(j.BaseFloder(), "http")
}

func (j *JinApp) ConsoleFolder() string {
	return filepath.Join(j.BaseFloder(), "console")
}

func (j *JinApp) ProviderFloder() string {
	return filepath.Join(j.BaseFloder(), "provider")
}

func (j *JinApp) MiddlewareFloder() string {
	return filepath.Join(j.HttpFolder(), "middleware")
}

func (j *JinApp) CommandFloder() string {
	return filepath.Join(j.ConsoleFolder(), "command")
}

func (j *JinApp) RuntimeFloder() string {
	return filepath.Join(j.StorageFolder(), "runtime")
}

func (j *JinApp) TestFloder() string {
	return filepath.Join(j.BaseFloder(), "test")
}
