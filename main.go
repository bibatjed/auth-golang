package main

import (
	"auth-golang/db"
	"auth-golang/routes"
	"github.com/gin-gonic/gin"
)

func init() {
}
func main() {
	initDB.InitializeDB()
	router := gin.Default()
	auth.Route(router)
	router.Run()
}
