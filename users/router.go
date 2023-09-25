package users

import (
	"github.com/gin-gonic/gin"
)


func AddUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users");
	users.GET("/",func(c *gin.Context) {
		go List(c)
	})

	/* users.POST("/", func(c *gin.Context) {
		createCh := make(chan gin.H)
		go CreateOne(c, createCh)
		c.JSON(http.StatusCreated, <-createCh)
	}) */

	users.GET("/:id",func(c *gin.Context) {
		go GetOne(c)
	})

	/* users.PUT("/:id",func(c *gin.Context) {
		updateOneCh := make(chan gin.H)
		go UpdateOne(c, updateOneCh)
		c.JSON(http.StatusOK, <-updateOneCh)
	})

	users.DELETE("/:id",func(c *gin.Context) {
		deleteOneCh := make(chan gin.H)
		go DeleteOne(c, deleteOneCh)
		c.JSON(http.StatusOK, <-deleteOneCh)
	}) */
}