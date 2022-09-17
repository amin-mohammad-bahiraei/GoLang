package Controllers

import (
	"GoLang/Data"
	"GoLang/Middleware"
	"GoLang/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProductList(c *gin.Context) {
	if err := Middleware.IsAuthorized(c); err != nil {
		return
	}

	db := Data.Database
	var products []Models.Product
	err := db.Find(&products).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, Models.Response{
			Message:   err,
			ErrorCode: 1,
		})
		return
	}

	c.JSON(http.StatusOK, Models.Response{
		Data: products,
	})
}
