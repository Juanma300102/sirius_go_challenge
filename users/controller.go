package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	listCh := make(chan queryResult)
	go ListAll(listCh)
	result := <- listCh
	if result.Err != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Err)
	}
	c.JSON(http.StatusOK, gin.H{"results": result.Result})
}

func GetOne(c *gin.Context) {
	getOneCh := make(chan queryResult)
	var parameters detailUriParameters
	if err := c.ShouldBindUri(&parameters); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	go Retrieve(parameters.Id, getOneCh)
	result := <- getOneCh
	if result.Err != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Err)
	}
	c.JSON(http.StatusOK, gin.H{"result": result.Result})
}

/* func CreateOne(c *gin.Context, ch chan gin.H) {
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
} */

/* func UpdateOne(c *gin.Context, ch chan gin.H) {
	var parameters detailUriParameters
	var updateDto updateUserDto

	if err := c.ShouldBindUri(&parameters); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	if err := c.ShouldBindJSON(&updateDto); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	} 

	user, err := Update(parameters.Id, &updateDto)
	
	if err!= nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	response := gin.H{"result": user}
	ch <- response
} */

/* func DeleteOne(c *gin.Context, ch chan gin.H) {
	var parameters detailUriParameters
	if err := c.ShouldBindUri(&parameters); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	user, err := Delete(parameters.Id)
	if err!= nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	response := gin.H{"result": user}
	ch <- response   
} */