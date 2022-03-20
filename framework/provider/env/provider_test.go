package env

import (
	"github.com/lai416703504/jin/framework"
	"github.com/lai416703504/jin/framework/contract"
	"github.com/lai416703504/jin/framework/provider/app"
	tests "github.com/lai416703504/jin/test"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestHadeEnvProvider(t *testing.T) {
	Convey("test hade env normal case", t, func() {
		basePath := tests.BasePath
		c := framework.NewJinContainer()
		sp := &app.JinAppProvider{BaseFolder: basePath}

		err := c.Bind(sp)
		So(err, ShouldBeNil)

		sp2 := &JinEnvProvider{}
		err = c.Bind(sp2)
		So(err, ShouldBeNil)

		envServ := c.MustMake(contract.EnvKey).(contract.Env)
		So(envServ.AppEnv(), ShouldEqual, "development")
		// So(envServ.Get("DB_HOST"), ShouldEqual, "127.0.0.1")
		// So(envServ.AppDebug(), ShouldBeTrue)
	})
}
