package private

import (
	initDB "auth-golang/db"
	"auth-golang/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProfile(c *gin.Context) {
	var user models.User
	v, _ := c.Get("user")
	initDB.Database.First(&user, v)

	c.JSON(http.StatusOK, gin.H{"message": "success", "result": user.ToClean()})
}
