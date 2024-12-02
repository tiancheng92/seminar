package controllers

import (
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/tiancheng92/seminar/controllers/api"
	"github.com/tiancheng92/seminar/controllers/api/universal"
	ginplus "github.com/tiancheng92/seminar/pkg/gin-plus"
	"github.com/tiancheng92/seminar/pkg/http/middleware/cross_domain"
	"github.com/tiancheng92/seminar/pkg/http/middleware/handle_error"
	"github.com/tiancheng92/seminar/pkg/log"
	"time"
)

// InitRouter 初始化路由
func InitRouter() *ginplus.Engine {
	r := ginplus.New()

	r.Use(
		ginzap.GinzapWithConfig(log.GetLogger(), &ginzap.Config{
			TimeFormat: time.DateTime,
			UTC:        true,
			SkipPaths:  []string{"/healthz"},
		}),
		gin.Recovery(),
		handle_error.HandleError,
		crossDomain.CrossDomain(),
	)

	r.NoRoute(universal.NoRoute)

	r.GET("healthz", universal.HealthCheck)

	apiGroup := r.Group("api")
	{
		api.NewExampleSceneRouter(apiGroup)
	}
	return r
}
