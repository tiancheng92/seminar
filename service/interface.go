package service

import (
	ginplus "github.com/tiancheng92/seminar/pkg/gin-plus"
	"github.com/tiancheng92/seminar/store/model"
	"github.com/tiancheng92/seminar/types/paginate"
)

type GenericInterface[M model.Interface] interface {
	ReadOnlyGenericInterface[M]
	Update(ctx *ginplus.Context, pk any, request any) (*M, error)
	Create(ctx *ginplus.Context, request any) (*M, error)
	Delete(ctx *ginplus.Context, pk any) error
}

type ReadOnlyGenericInterface[M model.Interface] interface {
	Get(ctx *ginplus.Context, pk any) (*M, error)
	List(ctx *ginplus.Context, pq *paginate.Query) (*paginate.Data[M], error)
	All(ctx *ginplus.Context) ([]*M, error)
	Distinct(ctx *ginplus.Context, field string) ([]string, error)
}
