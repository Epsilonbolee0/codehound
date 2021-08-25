package domain

import "github.com/lib/pq"

type Test struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	InArgs      pq.StringArray `json:"in_args" gorm:"type:text[]"`
	OutArgs     pq.StringArray `json:"out_args" gorm:"type:text[]"`
	Description string         `json:"description"`

	Implementation string `json:"implementation"`
	Author         uint   `json:"author"`
}

type TestDTO struct {
	ID             uint   `json:"id"`
	Login          string `json:"login"`
	Description    string `json:"description"`
	Implementation string `json:"implementation"`
	Author         uint   `json:"author"`
	Index          uint   `json:"index"`
	Value          string `json:"value"`
}
