package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tiancheng92/seminar/config"
	"github.com/tiancheng92/seminar/controllers"
)

func Run() {
	serverConfig := config.GetConf().Server
	switch serverConfig.Mode {
	case "debug":
		gin.SetMode(gin.DebugMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	s := &http.Server{
		Addr:    serverConfig.Host,        // 监听地址
		Handler: controllers.InitRouter(), // 处理器
	}

	go s.ListenAndServe()

	quit := make(chan os.Signal)
	defer close(quit)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s.Shutdown(ctx)
}
