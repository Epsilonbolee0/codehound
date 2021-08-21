package domain

type Tree struct {
	Name       string  `json:"name" gorm:"primaryKey"`
	Link       string  `json:"link"`
	ParentName *string `json:"parent_name"`

	Version Version `json:"version" gorm:"foreignKey:name"`
	Parent  *Tree   `json:"parent"`
}

type TreeMinimized struct {
	Name       string  `json:"name"`
	Link       string  `json:"link"`
	ParentName *string `json:"parent_name"`
}
