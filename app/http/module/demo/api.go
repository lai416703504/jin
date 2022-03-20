package demo

import (
	demoService "github.com/lai416703504/jin/app/provider/demo"
	"github.com/lai416703504/jin/framework/contract"
	"github.com/lai416703504/jin/framework/gin"
	"net/http"
)

// todo
type DemoApi struct {
	service *Service
}

func NewDemoApi() *DemoApi {
	service := NewService()
	return &DemoApi{
		service: service,
	}
}

func Register(r *gin.Engine) error {
	api := NewDemoApi()
	r.Bind(&demoService.DemoProvider{})

	r.GET("/demo/demo", api.Demo)
	r.GET("/demo/demo2", api.Demo2)
	r.POST("/demo/demo_post", api.DemoPost)

	return nil
}

//Demo godoc
// @Summary 获取所有用户
// @Description 获取所有用户
// @Produce json
// @Tags demo
// @Success 200 array []UserDTO
// @Router /demo/demo [get]
func (api *DemoApi) Demo(ctx *gin.Context) {
	//获取password
	configService := ctx.MustMake(contract.ConfigKey).(contract.Config)
	password := configService.GetString("database.mysql.password")
	//打印出来
	ctx.JSON(http.StatusOK, password)
	//appService := ctx.MustMake(contract.AppKey).(contract.App)
	//baseFolder := appService.BaseFolder()
	//ctx.JSON(200, baseFolder)

	//users := api.service.GetUsers()
	//usersDTO := UserModelsToUserDTOs(users)
	//c.JSON(200, usersDTO)

}

// Demo godoc
// @Summary 获取所有学生
// @Description 获取所有学生
// @Produce  json
// @Tags demo
// @Success 200 array []UserDTO
// @Router /demo/demo2 [get]
func (api *DemoApi) Demo2(ctx *gin.Context) {
	demoProvider := ctx.MustMake(demoService.DemoKey).(demoService.IService)
	students := demoProvider.GetAllStudent()
	usersDTO := StudentsToUserDTOs(students)
	ctx.JSON(200, usersDTO)
}

func (api *DemoApi) DemoPost(ctx *gin.Context) {
	type Foo struct {
		Name string
	}
	foo := &Foo{}
	err := ctx.BindJSON(&foo)
	if err != nil {
		ctx.AbortWithError(500, err)
	}
	ctx.JSON(200, nil)
}
