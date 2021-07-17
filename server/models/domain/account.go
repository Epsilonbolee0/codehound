package domain

type Account struct {
	UserID   uint   `json:"id" gorm:"primaryKey"`
	Login    string `json:"login" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role_id"`
}
