package domain

import "time"

type Version struct {
	Name  string    `json:"name" gorm:"primaryKey"`
	Code  string    `json:"code"`
	Title string    `json:"title"`
	Date  time.Time `json:"date"`

	Author     uint      `json:"author"`
	LanguageID uint      `json:"language"`
	Libraries  []Library `json:"libraries" gorm:"many2many:version_libraries"`
}

type VersionDTO struct {
	Name            string `json:"name"`
	Code            string `json:"code"`
	Title           string `json:"title"`
	Login           string `json:"login"`
	LanguageName    string `json:"language_name"`
	LanguageVersion string `json:"language_version"`
	LibraryName     string `json:"library_name"`
	LibraryVersion  string `json:"library_version"`
}
