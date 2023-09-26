package middlewares

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ErrorHandler() gin.HandlerFunc {
	return func (c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			errorsList := []string{}
			for _, err := range c.Errors {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					c.JSON(-1, gin.H{"Error": err.Error()})
					return
				}
				errorsList = append(errorsList, err.Error())
			}
			c.JSON(-1, gin.H{"Errors": errorsList})
		}
	}
}