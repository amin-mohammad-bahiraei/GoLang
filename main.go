package main

import (
	"GoLang/controller"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	initController := controller.InitController()
	initController.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	initController.Run()
}
