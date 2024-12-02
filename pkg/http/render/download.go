package render

import (
	"bytes"
	"strconv"

	"github.com/Yostardev/gf"
	"github.com/gin-gonic/gin"
)

func ResponseDownloadSteam(ctx *gin.Context, buf *bytes.Buffer, fileName string) {
	if ctx.IsAborted() {
		return
	}
	ctx.Abort()

	ctx.Header("Content-Length", strconv.Itoa(buf.Len()))
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Content-Disposition", gf.StringJoin("attachment;filename=", fileName))
	p := make([]byte, 4096)

	for {
		count, err := buf.Read(p)
		if err != nil {
			break
		}
		if count > 0 {
			if _, err := ctx.Writer.Write(p[:count]); err != nil {
				break
			}
			ctx.Writer.Flush()
		}
	}
}
