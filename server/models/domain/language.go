package domain

type Language struct {
	ID        uint      `json:"language_id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"index:lang_name_and_version,unique"`
	Version   string    `json:"version" gorm:"index:lang_name_and_version,unique"`
	Libraries []Library `json:"libraries" gorm:"foreignKey:LanguageID"`
}
