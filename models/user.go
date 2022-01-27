package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}