package service

import (
	ginplus "github.com/tiancheng92/seminar/pkg/gin-plus"
	"github.com/tiancheng92/seminar/store/model"
	"github.com/tiancheng92/seminar/store/repository"
	"github.com/tiancheng92/seminar/types/paginate"
	"github.com/tiancheng92/seminar/types/request"
	"reflect"
)

type genericService[M model.Interface] struct {
	*readOnlyGenericService[M]
	repo repository.GenericInterface[M]
}

func newGenericService[M model.Interface](repo repository.GenericInterface[M]) *genericService[M] {
	return &genericService[M]{
		newReadOnlyGenericService[M](repo),
		repo,
	}
}

func toModel[M model.Interface](requestPtr any) M {
	switch d := requestPtr.(type) {
	case request.Interface:
		return d.FormatToModel().(M)
	default:
		modelPtr := new(M)
		r := reflect.ValueOf(requestPtr).Elem()
		m := reflect.ValueOf(modelPtr).Elem()

		if r.Kind() != reflect.Struct || m.Kind() != reflect.Struct {
			panic("Both model and request input must be structs ptr")
		}

		for i := range r.NumField() {
			requestField := r.Type().Field(i)
			modelFieldValue := m.FieldByName(requestField.Name)
			if modelFieldValue.IsValid() && requestField.Type == modelFieldValue.Type() {
				modelFieldValue.Set(r.Field(i))
			}
		}
		return *modelPtr
	}
}

func (s *genericService[M]) Create(_ *ginplus.Context, request any) (*M, error) {
	return s.repo.Create(toModel[M](request))
}

func (s *genericService[M]) Update(_ *ginplus.Context, pk any, request any) (*M, error) {
	return s.repo.Update(pk, toModel[M](request))
}

func (s *genericService[M]) Delete(_ *ginplus.Context, pk any) error {
	return s.repo.Delete(pk)
}

// ------------------------------

type readOnlyGenericService[M model.Interface] struct {
	repo repository.GenericInterface[M]
}

func newReadOnlyGenericService[M model.Interface](repo repository.GenericInterface[M]) *readOnlyGenericService[M] {
	return &readOnlyGenericService[M]{repo}
}

func (ros *readOnlyGenericService[M]) Get(_ *ginplus.Context, pk any) (*M, error) {
	return ros.repo.Get(pk)
}

func (ros *readOnlyGenericService[M]) List(_ *ginplus.Context, pq *paginate.Query) (*paginate.Data[M], error) {
	return ros.repo.List(pq)
}

func (ros *readOnlyGenericService[M]) All(ctx *ginplus.Context) ([]*M, error) {
	return ros.repo.All(ctx.Request.URL.Query())
}
