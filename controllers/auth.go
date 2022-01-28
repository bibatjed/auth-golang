package authController

import (
	"auth-golang/db"
	"auth-golang/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {

}

func SignUp(c *gin.Context) {

	var requestBody models.User
	var user models.User
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	initDB.Database.Find(&user, "email = ?", requestBody.Email)
	if user.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	rows := initDB.Database.Create(&requestBody)

	if rows.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Something wrong"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}
