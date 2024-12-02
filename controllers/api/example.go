package api

import (
	ginplus "github.com/tiancheng92/seminar/pkg/gin-plus"
	"github.com/tiancheng92/seminar/service"
	"github.com/tiancheng92/seminar/store/model"
	"github.com/tiancheng92/seminar/types/request"
)

type exampleSceneController struct {
	*genericController[request.Example, model.Example]
}

func NewExampleSceneRouter(group *ginplus.RouterGroup) {
	c := &exampleSceneController{newGenericController[request.Example, model.Example](service.NewExampleService())}
	g := group.Group("example")
	{
		g.GET(":pk", c.Get)
		g.GET("", c.List)
		g.GET("all", c.All)
		g.POST("", c.Create)
		g.PUT(":pk", c.Update)
		g.DELETE(":pk", c.Delete)
	}
}
