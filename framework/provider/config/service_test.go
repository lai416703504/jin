package config

import (
	"github.com/lai416703504/jin/framework/contract"
	tests "github.com/lai416703504/jin/test"
	. "github.com/smartystreets/goconvey/convey"
	"path/filepath"
	"testing"
)

func TestHadeConfig_GetInt(t *testing.T) {
	Convey("test hade env normal case", t, func() {
		basePath := tests.BasePath
		folder := filepath.Join(basePath, "config")
		serv, err := NewJinConfig(folder, map[string]string{}, contract.EnvDevelopment)
		So(err, ShouldBeNil)
		conf := serv.(*JinConfig)
		timeout := conf.GetInt("database.mysql.timeout")
		So(timeout, ShouldEqual, 1)
	})
}
