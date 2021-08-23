package domain

type Tag struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Category string `json:"category"`
	Content  string `json:"content"`

	Versions []Version `json:"version" gorm:"many2many:version_tags"`
}
