package router

import (
	"challenge/users"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

// Run will start the server
func Run() {
	getRoutes()
	router.Run(":8080")
}

func getRoutes() {
	v1 := router.Group("/api/v1")
	users.AddUserRoutes(v1)
}