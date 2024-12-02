package universal

import (
	"github.com/gin-gonic/gin"
	"github.com/tiancheng92/seminar/pkg/errors"
	"github.com/tiancheng92/seminar/pkg/errors/ecode"
	"github.com/tiancheng92/seminar/pkg/http/render"
)

func NoRoute(ctx *gin.Context) {
	render.Response(ctx, nil, errors.WithCode(ecode.ErrPageNotFound, "Page not found"))
}
