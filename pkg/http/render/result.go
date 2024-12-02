package render

import (
	"github.com/Yostardev/json"
	"github.com/tiancheng92/seminar/pkg/errors"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/tiancheng92/seminar/types/paginate"
)

// PaginateInterface 定义分页接口
type PaginateInterface interface {
	GetPaginate() *paginate.Info
	GetItems() any
}

// Result 表示 HTTP 响应的结构体
type Result struct {
	Data any    `json:"data"` // 返回数据/错误详细信息
	Msg  string `json:"msg"`  // 请求结果
	Code int    `json:"code"` // 状态码
}

// reset 重置 Result 对象
func (r *Result) reset() {
	r.Data = nil
	r.Msg = "Success"
	r.Code = 200
}

// result结果池
var resultPool = &sync.Pool{
	New: func() any {
		return &Result{
			Msg:  "Success",
			Code: 200,
		}
	},
}

// PaginateData 表示分页数据的结构体
type PaginateData struct {
	Items    any            `json:"items"`    // 数据详情列表
	Paginate *paginate.Info `json:"paginate"` // 分页信息
}

// reset 重置 PaginateData 对象
func (r *PaginateData) reset() {
	r.Items = nil
	r.Paginate = new(paginate.Info)
}

// 分页数据池
var paginateDataPool = &sync.Pool{
	New: func() any {
		return &PaginateData{
			Paginate: new(paginate.Info),
		}
	},
}

// Response 处理并返回 HTTP 响应
func Response(ctx *gin.Context, data any, err error) {
	if ctx.IsAborted() {
		return
	}
	ctx.Abort()

	result := resultPool.Get().(*Result)
	defer func() {
		result.reset()
		resultPool.Put(result)
	}()

	if err != nil {
		handleError(ctx, result, err)
	} else {
		handleSuccess(ctx, result, data)
	}
}

// handleError 处理错误并设置响应
func handleError(ctx *gin.Context, result *Result, err error) {
	coder := errors.ParseCoder(err)
	result.Code = coder.HTTPStatus()
	if coder.String() != "" {
		result.Msg = coder.String()
	} else {
		result.Msg = err.Error()
	}

	if result.Code >= 400 && result.Code < 500 {
		ctx.Set("warn", err)
	} else {
		ctx.Set("error", err)
	}
	result.Data = err.Error()
	b, _ := json.Marshal(result)
	ctx.Data(result.Code, "application/json", b)
}

// handleSuccess 处理成功响应
func handleSuccess(ctx *gin.Context, result *Result, data any) {
	if d, ok := data.(PaginateInterface); ok {
		pd := paginateDataPool.Get().(*PaginateData)
		defer func() {
			pd.reset()
			paginateDataPool.Put(pd)
		}()
		pd.Items = d.GetItems()
		pd.Paginate = d.GetPaginate()
		result.Data = pd
	} else {
		result.Data = data
	}
	b, _ := json.Marshal(result)
	ctx.Data(200, "application/json", b)
}
