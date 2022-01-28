package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginFields struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	ID uint `json:"id" gorm:"primaryKey"`
	LoginFields
	Name string `json:"name" binding:"required"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	u.Password = string(hashedPassword)
	return
}

func (u *User) VerifyPassword(password string) bool {
	// Hashing the password with the default cost of 10
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return false
	}

	return true
}
