package handle_error

import (
	"github.com/gin-gonic/gin"
	"github.com/tiancheng92/seminar/pkg/log"
)

// HandleError 错误处理
func HandleError(ctx *gin.Context) {
	ctx.Next()
	if warnInfo, exist := ctx.Get("warn"); exist {
		log.Warnf("%v", warnInfo)
	}
	if errInfo, exist := ctx.Get("error"); exist {
		log.Errorf("%+v", errInfo)
	}
}
