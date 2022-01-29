package middleware

import (
	"auth-golang/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware(c *gin.Context) {
	bearer := c.GetHeader("authorization")

	if bearer == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	token := strings.Split(bearer, " ")

	if len(token) != 2 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bearer token not in proper format"})
		return
	}

	verifyToken, err := utils.VerifyToken(token[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	c.Set("user", verifyToken)
	c.Next()
}
