package domain

type Library struct {
	ID      uint   `json:"library_id" gorm:"primaryKey"`
	Name    string `json:"name" gorm:"index:lib_name_and_version,unique"`
	Version string `json:"version" gorm:"index:lib_name_and_version,unique"`

	LanguageID uint      `json:"language_id"`
	Versions   []Version `json:"versions" gorm:"many2many:version_libraries"`
}

type LibraryDTO struct {
	Name            string `json:"name"`
	Version         string `json:"version"`
	LanguageName    string `json:"language_name"`
	LanguageVersion string `json:"language_version"`
}
