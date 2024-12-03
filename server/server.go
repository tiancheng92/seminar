package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tiancheng92/seminar/config"
	"github.com/tiancheng92/seminar/controllers"
	"github.com/tiancheng92/seminar/pkg/log"
	"net/http"
	"os"
	"os/signal"
	"time"
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

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Fatalf("%+v", err)
		}
	}()

	quit := make(chan os.Signal)
	defer close(quit)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s.Shutdown(ctx)
}
