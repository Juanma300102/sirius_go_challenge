package users

import (
	"challenge/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context, ch chan gin.H) {
	users, err := ListAll()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	response := gin.H{"results": users}
	ch <- response

}

type getOneUriParameters struct {
	Id int `uri:"id" binding:"required"`
}

func GetOne(c *gin.Context, ch chan gin.H) {
	var parameters getOneUriParameters
	if err := c.ShouldBindUri(&parameters); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	user; err := Retrieve(parameters.Id)
	response := gin.H{"result": user}
	ch <- response   
}

func CreateOne(c *gin.Context, ch chan gin.H) {
	var createDto models.User
	c.BindJSON(&createDto)
	user, err := Create(&createDto)
	if err!= nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	response := gin.H{"result": user}
	ch <- response
}

func UpdateOne(c *gin.Context, ch chan gin.H) {
	response := gin.H{"message": "OK"}
	ch <- response
}

func DeleteOne(c *gin.Context, ch chan gin.H) {
	response := gin.H{"message": "OK"}
	ch <- response
}