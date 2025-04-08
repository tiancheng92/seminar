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

type Example struct {
	Name     string `json:"name" binding:"required"`
	Describe string `json:"describe" binding:"required"`
	Date     string `json:"date" binding:"required,date_time_format=2006-01-02"`
	Number   int    `json:"number" binding:"required"`
}

type Distinct struct {
	Field string `uri:"field" binding:"required"`
}
