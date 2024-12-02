package model

import (
	"database/sql/driver"
	"github.com/Yostardev/gf"
	"github.com/Yostardev/json"
	"github.com/tiancheng92/seminar/pkg/errors"
)

type Example struct {
	Model
	Name          string        `json:"name" gorm:"type:varchar(50);not null"`
	Describe      string        `json:"describe" gorm:"type:text;not null"`
	Date          string        `json:"date" gorm:"type:varchar(10);not null"`
	Number        int           `json:"number" gorm:"type:int;not null"`
	JsonFieldList JsonFieldList `json:"json_field_list" gorm:"type:json;not null"`
}

type JsonFieldList []*JsonField

type JsonField struct {
	FieldOne string `json:"field_one"`
	FieldTwo string `json:"field_two"`
}

func (l *JsonFieldList) Scan(src any) error {
	switch v := src.(type) {
	case []byte:
		return json.Unmarshal(v, l)
	case string:
		return json.Unmarshal(gf.StringToBytes(v), l)
	default:
		return errors.New("invalid src type")
	}
}

func (l JsonFieldList) Value() (driver.Value, error) {
	return json.Marshal(&l)
}
