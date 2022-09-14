package Controllers

import (
	"GoLang/Api"
	"GoLang/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

var Database *gorm.DB

func init() {
	Db = Database
	router := Api.Router
	router.GET("/GetListFoo", GetListFoo)
	router.GET("/GetFooById/:id", GetFooById)
}

func GetListFoo(c *gin.Context) {
	foo := make([]Models.Foo, 0)
	err := Db.Find(&foo).Error
	if err != nil {
		c.JSON(http.StatusOK, Models.Response{
			Message:   err,
			ErrorCode: 1,
			Data:      nil,
		})
		return
	}
	c.JSON(http.StatusOK, Models.Response{
		Message:   "",
		ErrorCode: 0,
		Data:      foo,
	})
}

func GetFooById(c *gin.Context) {
	foo := Models.Foo{}
	id := c.Param("id")
	err := Db.Where("id = ?", id).First(&foo).Error
	if err != nil {
		c.JSON(http.StatusOK, Models.Response{
			Message:   err,
			ErrorCode: 1,
			Data:      nil,
		})
		return
	}
	c.JSON(http.StatusOK, Models.Response{
		Message:   "",
		ErrorCode: 0,
		Data:      foo,
	})
}
