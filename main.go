package main

import (
	"auth-golang/db"
	"auth-golang/routes/auth"
	"auth-golang/routes/private"
	"github.com/gin-gonic/gin"
)

func init() {
}
func main() {
	initDB.InitializeDB()
	router := gin.Default()
	auth.Route(router)
	private.Route(router)
	router.Run()
}
