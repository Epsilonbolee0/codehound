package domain

import "time"

type Version struct {
	Name  string    `json:"name" gorm:"primaryKey"`
	Code  string    `json:"code"`
	Title string    `json:"title"`
	Date  time.Time `json:"date"`

	Author uint `json:"author"`
}

type VersionDTO struct {
	Name  string `json:"name"`
	Code  string `json:"code"`
	Title string `json:"title"`
	Login string `json:"login"`
}
