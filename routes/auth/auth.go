package auth

import (
	authController "auth-golang/controllers/auth"
	"github.com/gin-gonic/gin"
)

func Route(route *gin.Engine) {
	auth := route.Group("/auth")
	{
		auth.POST("/login", authController.Login)
		auth.POST("/sign-up", authController.SignUp)
	}
}
