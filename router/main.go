package router

import (
	"challenge/db"
	"challenge/middlewares"
	"challenge/users"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	getRoutes(r)
	return r
}

// Run will start the server
func Run() {
	r := SetupRouter()
	db.Migrate()
	r.Run(":8080")
}

func getRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	v1.Use(middlewares.ErrorHandler())
	users.AddUserRoutes(v1)
}