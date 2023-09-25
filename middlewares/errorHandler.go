package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ErrorHandler(c *gin.Context) gin.HandlerFunc {
	return func (c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {
			switch err.Err {
				case gorm.ErrRecordNotFound:
					fmt.Println(err.Err.Error())
				c.JSON(http.StatusNotFound, gin.H{"error": err.Err.Error()})  
			}
		}
		fmt.Println("EN HANDLER")

		c.JSON(http.StatusInternalServerError, "")
	}
}