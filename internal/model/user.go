package model

type User struct {
	ID       int64  `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"unique"`
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Status   int8   `json:"status"`
}

type RegisterRequest struct {
	// Username string `json:"username" binding:"required,min=3,max=20"`
	// Age      uint8  `json:"age" binding:"required,gte=1,lte=130"`
	// Gender   uint8  `json:"gender" binding:"oneof=0 1"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserQuery struct {
	Page  int    `form:"page" binding:"omitempty,gte=1"`
	Email string `form:"email" binding:"omitempty,email"`
}
