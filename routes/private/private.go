package private

import (
	privateController "auth-golang/controllers/private"
	"auth-golang/middleware"
	"github.com/gin-gonic/gin"
)

func Route(route *gin.Engine) {
	private := route.Group("/private")
	{
		private.GET("/", middleware.AuthMiddleware, privateController.GetProfile)
	}
}
