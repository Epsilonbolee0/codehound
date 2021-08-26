package domain

type Description struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	Content        string `json:"content" gorm:"unique"`
	Author         uint   `json:"author"`
	Implementation string `json:"implementation"`
}

type DescriptionDTO struct {
	ID             uint   `json:"id"`
	Implementation string `json:"implementation"`
	Content        string `json:"content"`
}
