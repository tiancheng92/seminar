package repository

import (
	"github.com/tiancheng92/seminar/store/model"
	"github.com/tiancheng92/seminar/types/paginate"
	"net/url"
)

type GenericInterface[M model.Interface] interface {
	Get(pk any) (*M, error)
	Create(attributes M) (*M, error)
	Update(pk any, attributes M) (*M, error)
	Delete(pk any) error
	List(pq *paginate.Query) (*paginate.Data[M], error)
	Count(params url.Values) (int64, error)
	All(params url.Values) ([]*M, error)
	Distinct(field string) ([]string, error)
}

type ExampleRepoInterface interface {
	GenericInterface[model.Example]
}
