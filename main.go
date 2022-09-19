package main

import (
	"GoLang/Api"
	"fmt"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println(err)
	}

	api := Api.Api
	if err := api.Run(); err != nil {
		fmt.Println(err)
	}
}
