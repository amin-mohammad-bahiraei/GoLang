package controller

import (
	"GoLang/controller/api"
	"github.com/gin-gonic/gin"
)

func InitController() *gin.Engine {
	r := gin.New()
	apiGroup := r.Group("/api")
	apiGroup.GET("/GetListFoo", api.GetListFoo)
	apiGroup.GET("/GetFooById/:id", api.GetFooById)

	apiGroup.GET("/auth/login", api.Login)
	return r
}
