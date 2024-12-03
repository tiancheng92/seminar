package gin_plus

import (
	"github.com/gin-gonic/gin"
	"github.com/tiancheng92/seminar/pkg/errors"
	"github.com/tiancheng92/seminar/pkg/http/render"
	"github.com/tiancheng92/seminar/pkg/validator"
	"github.com/tiancheng92/seminar/types/paginate"
	"strconv"
)

type Engine struct {
	*gin.Engine
}

func New() *Engine {
	return &Engine{gin.New()}
}

func (e *Engine) Group(relativePath string, handlers ...gin.HandlerFunc) *RouterGroup {
	return &RouterGroup{e.Engine.Group(relativePath, handlers...)}
}

type Context struct {
	*gin.Context
}

// do 执行函数，如果上下文没有被中止
func (c *Context) do(f func()) *Context {
	if !c.IsAborted() {
		f()
	}
	return c
}

// HandleAndRender 处理函数并渲染响应
func (c *Context) HandleAndRender(f any) {
	c.do(func() {
		var err error
		var resp any

		// 根据不同的函数类型处理
		switch fn := f.(type) {
		case func():
			fn()
		case func() error:
			err = fn()
		case func() (any, error):
			resp, err = fn()
		default:
			err = errors.New("invalid function type")
		}
		// 渲染响应
		render.Response(c.Context, resp, err)
	})
}

// BindBody 绑定请求体到指定结构体
func (c *Context) BindBody(ptr any) *Context {
	return c.do(func() {
		if err := c.ShouldBind(ptr); err != nil {
			render.Response(c.Context, nil, validator.HandleValidationErr(err))
		}
	})
}

// BindQuery 绑定查询参数到指定结构体
func (c *Context) BindQuery(ptr any) *Context {
	return c.do(func() {
		if err := c.ShouldBindQuery(ptr); err != nil {
			render.Response(c.Context, nil, validator.HandleValidationErr(err))
		}
	})
}

// BindParams 绑定 URI 参数到指定结构体
func (c *Context) BindParams(ptr any) *Context {
	return c.do(func() {
		if err := c.ShouldBindUri(ptr); err != nil {
			render.Response(c.Context, nil, validator.HandleValidationErr(err))
		}
	})
}

// BindHeader 绑定请求头到指定结构体
func (c *Context) BindHeader(ptr any) *Context {
	return c.do(func() {
		if err := c.ShouldBindHeader(ptr); err != nil {
			render.Response(c.Context, nil, validator.HandleValidationErr(err))
		}
	})
}

// BindPaginateQuery 绑定分页查询参数
func (c *Context) BindPaginateQuery(ptr *paginate.Query) *Context {
	return c.do(func() {
		page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
		if err != nil || page < 1 {
			page = 1
		}

		pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", "20"))
		if err != nil || pageSize < 1 {
			pageSize = 20
		}

		ptr.Page = page
		ptr.PageSize = pageSize
		ptr.Order = c.DefaultQuery("order", "")
		ptr.OrderBy = c.DefaultQuery("order_by", "")
		ptr.Search = c.DefaultQuery("search", "")
		ptr.Params = c.Request.URL.Query()
	})
}

type RouterGroup struct {
	*gin.RouterGroup
}

func (g *RouterGroup) Group(relativePath string, handlers ...gin.HandlerFunc) *RouterGroup {
	return &RouterGroup{g.RouterGroup.Group(relativePath, handlers...)}
}

func (g *RouterGroup) GET(relativePath string, f func(*Context)) {
	g.RouterGroup.GET(relativePath, func(c *gin.Context) {
		f(&Context{c})
	})
}

func (g *RouterGroup) POST(relativePath string, f func(*Context)) {
	g.RouterGroup.POST(relativePath, func(c *gin.Context) {
		f(&Context{c})
	})
}

func (g *RouterGroup) PUT(relativePath string, f func(*Context)) {
	g.RouterGroup.PUT(relativePath, func(c *gin.Context) {
		f(&Context{c})
	})
}

func (g *RouterGroup) PATCH(relativePath string, f func(*Context)) {
	g.RouterGroup.PATCH(relativePath, func(c *gin.Context) {
		f(&Context{c})
	})
}

func (g *RouterGroup) DELETE(relativePath string, f func(*Context)) {
	g.RouterGroup.DELETE(relativePath, func(c *gin.Context) {
		f(&Context{c})
	})
}
