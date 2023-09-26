package users

import (
	"github.com/gin-gonic/gin"
)


func AddUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users");
	users.GET("/",func(c *gin.Context) {
		List(c)
	})

	users.POST("/", func(c *gin.Context) {
		CreateOne(c)
	})

	users.GET("/:id",func(c *gin.Context) {
		GetOne(c)
	})

	users.PUT("/:id",func(c *gin.Context) {
		UpdateOne(c)
	})

	users.DELETE("/:id",func(c *gin.Context) {
		DeleteOne(c)
	})
}