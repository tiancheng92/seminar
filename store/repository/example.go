package repository

import (
	"github.com/tiancheng92/seminar/pkg/errors"
	"github.com/tiancheng92/seminar/pkg/errors/ecode"
	"github.com/tiancheng92/seminar/store/model"
	"github.com/tiancheng92/seminar/types/paginate"
)

type exampleRepository struct {
	*genericRepository[model.Example] // 继承泛型实现
}

func NewExampleRepository() ExampleRepoInterface {
	return &exampleRepository{newGenericRepository[model.Example]()}
}

// List 复写泛形实现
func (r *exampleRepository) List(pq *paginate.Query) (*paginate.Data[model.Example], error) {
	r.paginateData.Init(pq)
	session := r.db.Model(new(model.Example)).Scopes(Paginate[model.Example](pq))
	if fieldOne, ok := pq.Params["json_field_list.field_one"]; ok {
		for i := range fieldOne {
			session = session.Where("JSON_CONTAINS(`json_field_list`, JSON_OBJECT('field_one', ?))", fieldOne[i])
		}
	}
	if fieldTwo, ok := pq.Params["json_field_list.field_two"]; ok {
		for i := range fieldTwo {
			session = session.Where("JSON_CONTAINS(`json_field_list`, JSON_OBJECT('field_two', ?))", fieldTwo[i])
		}
	}
	err := session.Find(&r.paginateData.Items).Offset(-1).Limit(-1).Count(&r.paginateData.Total).Error
	return r.paginateData, errors.WithCode(ecode.ErrGet, err)
}
