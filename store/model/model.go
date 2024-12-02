package model

import (
	"gorm.io/gorm"
	"time"
)

type Interface interface {
	GetPrimaryKeyName() string
	GetFuzzySearchFieldList() []string
	GetDefaultOrderBy() string
	GetDefaultOrder() string
}

type Model struct {
	ID        uint64    `json:"id" gorm:"primary_key;type:bigint unsigned;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}

func (Model) GetPrimaryKeyName() string {
	return "id"
}

func (Model) GetFuzzySearchFieldList() []string {
	return []string{}
}

func (m Model) GetDefaultOrderBy() string {
	return m.GetPrimaryKeyName()
}

func (Model) GetDefaultOrder() string {
	return "desc"
}

type SoftDeleteModel struct {
	Model
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
