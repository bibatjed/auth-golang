package authController

import (
	"auth-golang/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	ID       uint   `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

func Login(c *gin.Context) {

}

func SignUp(c *gin.Context) {

	var requestBody User
	var user User
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	initDB.Database.Find(&user, "email = ?", requestBody.Email)
	if user.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	initDB.Database.Create(&requestBody)
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}
