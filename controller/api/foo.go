package api

import (
	"GoLang/database"
	"GoLang/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetListFoo(c *gin.Context) {
	results := database.DB().First(&model.Foo{})
	if results != nil {
		fmt.Println(results)
	}
	foo := []model.Foo{
		{ID: 1, Title: "title_1", Price: 4, CreatedAt: "a", UpdatedAt: "d"},
		{ID: 3, Title: "title_2", Price: 3, CreatedAt: "b", UpdatedAt: "c"},
		{ID: 2, Title: "title_3", Price: 2, CreatedAt: "c", UpdatedAt: "b"},
		{ID: 4, Title: "title_4", Price: 1, CreatedAt: "d", UpdatedAt: "a"},
	}
	fmt.Println(foo)
	c.JSON(http.StatusOK, foo)
}

func GetFooById(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"message": "ok",
	})
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
