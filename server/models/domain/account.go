package domain

type Account struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Login    string `json:"login" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role"`

	Versions []Version `json:"versions" gorm:"foreignKey:Author"`
}

type AccountDTO struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
