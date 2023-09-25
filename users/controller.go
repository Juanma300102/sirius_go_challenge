package users

import (
	"fmt"
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

func GetOne(c *gin.Context, ch chan gin.H) {
	var parameters detailUriParameters
	if err := c.ShouldBindUri(&parameters); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	user, err := Retrieve(parameters.Id)
	if err!= nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	response := gin.H{"result": user}
	ch <- response   
}

func CreateOne(c *gin.Context, ch chan gin.H) {
	var createDto createUserDto
	if err := c.ShouldBindJSON(&createDto); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	user, err := Create(&createDto)
	if err!= nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	response := gin.H{"result": user}
	ch <- response
}

func UpdateOne(c *gin.Context, ch chan gin.H) {
	var parameters detailUriParameters
	var updateDto updateUserDto

	if err := c.ShouldBindUri(&parameters); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	if err := c.ShouldBindJSON(&updateDto); err != nil {
		fmt.Println(err)
		c.AbortWithError(http.StatusBadRequest, err)
	} 

	user, err := Update(parameters.Id, &updateDto)
	
	if err!= nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	response := gin.H{"result": user}
	ch <- response
}

func DeleteOne(c *gin.Context, ch chan gin.H) {
	response := gin.H{"message": "OK"}
	ch <- response
}