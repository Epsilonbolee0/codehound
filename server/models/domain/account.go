package domain

type Account struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Login    string `json:"login" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Icon     string `json:"icon"`
	Password string `json:"password"`
	Role     string `json:"role"`

	Versions     []Version     `json:"versions" gorm:"foreignKey:Author"`
	Descriptions []Description `json:"descriptions" gorm:"foreignKey:Author"`
	Tests        []Test        `json:"tests" gorm:"foreignKey:Author"`
}

type AccountAPI struct {
	ID       uint   `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Icon     string `json:"icon"`
	Role     string `json:"role"`
}

type AccountDTO struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Icon     string `json:"icon"`
	Role     string `json:"role"`
}
