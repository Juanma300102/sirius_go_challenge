package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func AddUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users");
	users.GET("/",func(c *gin.Context) {
		listCh := make(chan gin.H)
		go List(c, listCh)
		c.JSON(http.StatusOK, <-listCh)
	})

	users.GET("/:id",func(c *gin.Context) {
		getOneCh := make(chan gin.H)
		go GetOne(c, getOneCh)
		c.JSON(http.StatusOK, <-getOneCh)
	})

	users.POST("/", func(c *gin.Context) {
		createCh := make(chan gin.H)
		go CreateOne(c, createCh)
		c.JSON(http.StatusCreated, <-createCh)
	})
}