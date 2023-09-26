package users

import (
	"challenge/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func List(c *gin.Context) {
	listCh := make(chan queryResult)
	go ListAll(listCh)
	result := <- listCh
	if result.Err != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Err)
	}
	c.JSON(http.StatusOK, gin.H{"results": result.Result, "count": len(result.Result.([]models.User))})
}

func GetOne(c *gin.Context) {
	var parameters detailUriParameters
	if err := c.ShouldBindUri(&parameters); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	getOneCh := make(chan queryResult)
	go Retrieve(parameters.Id, getOneCh)
	result := <- getOneCh
	
	if result.Err != nil {
		if errors.Is(result.Err, gorm.ErrRecordNotFound){
			c.AbortWithError(http.StatusNotFound, result.Err)	
		} else {
			c.AbortWithError(http.StatusInternalServerError, result.Err)
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": result.Result})
}

func CreateOne(c *gin.Context) {
	var createDto createUserDto
	if err := c.ShouldBindJSON(&createDto); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	createCh := make(chan queryResult)
	go Create(&createDto, createCh)
	result := <- createCh
	if result.Err!= nil {
		if errors.Is(result.Err, gorm.ErrRecordNotFound){
			c.AbortWithError(http.StatusNotFound, result.Err)	
		} else {
			c.AbortWithError(http.StatusInternalServerError, result.Err)
		}
		return
	}
	c.JSON(http.StatusCreated, gin.H{"result": result.Result})
}

func UpdateOne(c *gin.Context) {
	var parameters detailUriParameters
	var updateDto updateUserDto

	if err := c.ShouldBindUri(&parameters); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := c.ShouldBindJSON(&updateDto); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	} 

	updateCh := make(chan queryResult)
	go Update(parameters.Id, &updateDto, updateCh)
	result := <- updateCh
	
	if result.Err!= nil {
		if errors.Is(result.Err, gorm.ErrRecordNotFound){
			c.AbortWithError(http.StatusNotFound, result.Err)	
		} else {
			c.AbortWithError(http.StatusInternalServerError, result.Err)
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": result.Result})
}

func DeleteOne(c *gin.Context) {
	var parameters detailUriParameters
	if err := c.ShouldBindUri(&parameters); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	deleteCh := make(chan queryResult)
	go Delete(parameters.Id, deleteCh)
	result := <- deleteCh
	if result.Err != nil {
		if errors.Is(result.Err, gorm.ErrRecordNotFound){
			c.AbortWithError(http.StatusNotFound, result.Err)	
		} else {
			c.AbortWithError(http.StatusInternalServerError, result.Err)
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": result.Result})
}