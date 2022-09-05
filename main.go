package main

import (
	"GoLang/controller"
)

func main() {
	initController := controller.InitController()
	initController.Run(":8080")
}
