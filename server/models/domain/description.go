package domain

type Description struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	Content        string `json:"content"`
	Author         uint   `json:"author"`
	Implementation string `json:"version"`
}
