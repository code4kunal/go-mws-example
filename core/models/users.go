package models

// User model
type UserLoginDetails struct {
	UserName   string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}


type UserRegistration struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserCity  string `json:"user_city"`
}
