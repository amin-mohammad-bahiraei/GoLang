package controller

import (
	"GoLang/controller/api"
	"github.com/gin-gonic/gin"
)

func InitController() *gin.Engine {
	r := gin.New()
	apiGroup := r.Group("/api")
	apiGroup.GET("/foo", api.GetListFoo)
	return r
}
