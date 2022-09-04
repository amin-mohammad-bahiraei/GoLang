package main

import (
	"GoLang/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
)

func main() {
	foo := []model.Foo{
		{ID: 1, Title: "title", Price: 1, CreatedAt: "asdf", UpdatedAt: "asdf"},
		{ID: 3, Title: "title", Price: 1, CreatedAt: "asdf", UpdatedAt: "asdf"},
		{ID: 2, Title: "title", Price: 1, CreatedAt: "asdf", UpdatedAt: "asdf"},
		{ID: 4, Title: "title", Price: 1, CreatedAt: "asdf", UpdatedAt: "asdf"},
	}
	router := gin.Default()

	sort.Slice(foo, func(i, j int) bool {
		return foo[i].ID < foo[j].ID
	})

	group := router.Group("/test")

	group.GET("/bar", func(c *gin.Context) {
		c.JSON(http.StatusOK, foo)
	})

	router.Run(":8080")
}
