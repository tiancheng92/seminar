package repository

import (
	"github.com/tiancheng92/seminar/store/model"
)

type exampleRepository struct {
	*genericRepository[model.Example] // 继承泛型实现
}

func NewExampleRepository() ExampleRepoInterface {
	return &exampleRepository{newGenericRepository[model.Example]()}
}
