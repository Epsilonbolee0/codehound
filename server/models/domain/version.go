package domain

import "time"

type Version struct {
	Name string    `json:"name" gorm:"primaryKey"`
	Code string    `json:"code"`
	Date time.Time `json:"date"`

	Author         uint      `json:"author"`
	Implementation string    `json:"implementation"`
	Libraries      []Library `json:"libraries" gorm:"many2many:version_libraries"`
}

type VersionDTO struct {
	Name           string `json:"name"`
	Code           string `json:"code"`
	Login          string `json:"login"`
	Link           string `json:"link"`
	Implementation string `json:"implementation"`
	LibraryName    string `json:"library_name"`
	LibraryVersion string `json:"library_version"`
}
