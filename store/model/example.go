package model

type Example struct {
	Model
	Name     string `json:"name" gorm:"type:varchar(50);not null"`
	Describe string `json:"describe" gorm:"type:text;not null"`
	Date     string `json:"date" gorm:"type:varchar(10);not null"`
	Number   int    `json:"number" gorm:"type:int;not null"`
}
