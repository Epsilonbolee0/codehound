package domain

import "github.com/lib/pq"

type Implementation struct {
	Name       string `json:"name" gorm:"primaryKey"`
	LanguageID uint   `json:"language"`

	Versions    []Version      `json:"versions" gorm:"foreignKey:Implementation"`
	Description []Description  `json:"description" gorm:"foreignKey:Implementation"`
	InArgs      pq.StringArray `json:"in_args" gorm:"type:text[]"`
	OutArgs     pq.StringArray `json:"out_args" gorm:"type:text[]"`
}

type ImplementationDTO struct {
	Name            string `json:"name"`
	LanguageName    string `json:"language_name"`
	LanguageVersion string `json:"language_version"`
	Value           string `json:"value"`
	Index           uint   `json:"index"`
}
