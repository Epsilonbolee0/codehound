package domain

type Implementation struct {
	Name       string    `json:"name" gorm:"primaryKey"`
	LanguageID uint      `json:"language"`
	Versions   []Version `json:"versions" gorm:"foreignKey:Implementation"`
}

type ImplementationDTO struct {
	Name            string `json:"name"`
	LanguageName    string `json:"language_name"`
	LanguageVersion string `json:"language_version"`
}
