package request

import (
	"github.com/tiancheng92/seminar/store/model"
)

type Interface interface {
	FormatToModel() model.Interface
}

type PrimaryKey struct {
	PrimaryKey uint64 `uri:"pk" binding:"required"`
}

type Header struct {
	Authorization string `header:"Authorization" binding:"required"`
}

type Distinct struct {
	Field string `uri:"field" binding:"required"`
}
