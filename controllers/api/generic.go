package api

import (
	ginplus "github.com/tiancheng92/seminar/pkg/gin-plus"
	"github.com/tiancheng92/seminar/service"
	"github.com/tiancheng92/seminar/store/model"
	"github.com/tiancheng92/seminar/types/paginate"
	"github.com/tiancheng92/seminar/types/request"
)

type genericController[R any, M model.Interface] struct {
	*readOnlyGenericController[M]
	service.GenericInterface[M]
}

func newGenericController[R any, M model.Interface](service service.GenericInterface[M]) *genericController[R, M] {
	return &genericController[R, M]{newReadOnlyGenericController[M](service), service}
}

func (c *genericController[R, M]) Get(ctx *ginplus.Context) {
	c.readOnlyGenericController.Get(ctx)
}

func (c *genericController[R, M]) List(ctx *ginplus.Context) {
	c.readOnlyGenericController.List(ctx)
}

func (c *genericController[R, M]) All(ctx *ginplus.Context) {
	c.readOnlyGenericController.All(ctx)
}

func (c *genericController[R, M]) Distinct(ctx *ginplus.Context) {
	c.readOnlyGenericController.Distinct(ctx)
}

func (c *genericController[R, M]) Create(ctx *ginplus.Context) {
	r := new(R)
	ctx.BindBody(r).HandleAndRender(func() (any, error) {
		return c.GenericInterface.Create(ctx, r)
	})
}

func (c *genericController[R, M]) Update(ctx *ginplus.Context) {
	p, r := new(request.PrimaryKey), new(R)
	ctx.BindParams(p).BindBody(r).HandleAndRender(func() (any, error) {
		return c.GenericInterface.Update(ctx, p.PrimaryKey, r)
	})
}

func (c *genericController[R, M]) Delete(ctx *ginplus.Context) {
	p := new(request.PrimaryKey)
	ctx.BindParams(p).HandleAndRender(func() error {
		return c.GenericInterface.Delete(ctx, p.PrimaryKey)
	})
}

// ------------------------------

type readOnlyGenericController[M model.Interface] struct {
	service.ReadOnlyGenericInterface[M]
}

func newReadOnlyGenericController[M model.Interface](service service.ReadOnlyGenericInterface[M]) *readOnlyGenericController[M] {
	return &readOnlyGenericController[M]{service}
}

func (roc *readOnlyGenericController[M]) Get(ctx *ginplus.Context) {
	p := new(request.PrimaryKey)
	ctx.BindParams(p).HandleAndRender(func() (any, error) {
		return roc.ReadOnlyGenericInterface.Get(ctx, p.PrimaryKey)
	})
}

func (roc *readOnlyGenericController[M]) List(ctx *ginplus.Context) {
	pq := new(paginate.Query)
	ctx.BindPaginateQuery(pq).HandleAndRender(func() (any, error) {
		return roc.ReadOnlyGenericInterface.List(ctx, pq)
	})
}

func (roc *readOnlyGenericController[M]) All(ctx *ginplus.Context) {
	ctx.HandleAndRender(func() (any, error) {
		return roc.ReadOnlyGenericInterface.All(ctx)
	})
}

func (roc *readOnlyGenericController[M]) Distinct(ctx *ginplus.Context) {
	p := new(request.Distinct)
	ctx.BindParams(p).HandleAndRender(func() (any, error) {
		return roc.ReadOnlyGenericInterface.Distinct(ctx, p.Field)
	})
}
