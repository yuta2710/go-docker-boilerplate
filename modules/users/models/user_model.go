package models

type InsertUserRequest struct {
	FirstName string `gorm:"column:firstName" json:"firstName"`
	LastName  string `gorm:"column:lastName" json:"lastName"`
	Email     string `gorm:"column:email" json:"email"`
	Password  string `gorm:"column:password" json:"password"`
	// Role      string `gorm:"column:role" json:"role"`
}

type GetUserByIdRequest struct {
	Id string `json:"id"`
}
