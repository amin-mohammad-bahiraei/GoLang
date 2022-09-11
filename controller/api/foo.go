package api

import (
	"GoLang/config"
	"GoLang/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetListFoo(c *gin.Context) {
	foo := make([]models.Foo, 0)
	if config.Database.Find(&foo).Error != nil {
		fmt.Println("Error Get List Foo")
	}
	c.JSON(http.StatusOK, foo)
}

func GetFooById(c *gin.Context) {
	foo := models.Foo{}
	id := c.Param("id")
	if config.Database.Where("id = ?", id).First(&foo).Error != nil {
		c.JSON(http.StatusOK, nil)
	} else {
		c.JSON(http.StatusOK, foo)
	}
}

func CreateFoo(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"message": "ok",
	})
}

func UpdateFoo(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"message": "ok",
	})
}

func DeleteFoo(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"message": "ok",
	})
}
