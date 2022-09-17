package Controllers

import (
	"GoLang/Data"
	"GoLang/Middleware"
	"GoLang/Models"
	"GoLang/services"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

var db *gorm.DB

func init() {
	db = Data.Database
}

func SignUp(c *gin.Context) {
	var user Models.User

	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, Models.Response{
			Message:   err,
			ErrorCode: 1,
			Data:      nil,
		})
		return
	}
	var dbUser Models.User
	db.Where("username = ? or phone_number = ?", user.Username, user.PhoneNumber).First(&dbUser)
	fmt.Println(dbUser)
	if dbUser.Username != "" || dbUser.PhoneNumber != "" {
		c.JSON(http.StatusBadRequest, Models.Response{
			Message:   "username is already exist",
			ErrorCode: 1,
			Data:      nil,
		})
		return
	}

	user.Password = Services.GenerateHash(user.Password)
	db.Create(&user)
	c.JSON(http.StatusBadRequest, Models.Response{
		Message:   "",
		ErrorCode: 0,
		Data:      user,
	})
}

func SignIn(c *gin.Context) {
	var login Models.Login
	err := json.NewDecoder(c.Request.Body).Decode(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, Models.Response{
			Message:   err,
			ErrorCode: 1,
			Data:      nil,
		})
		return
	}
	var user Models.User
	db.Where("username = ?", login.Username).First(&user)
	if user.Username == "" {
		c.JSON(http.StatusBadRequest, Models.Response{
			Message:   "username or password is incorrect",
			ErrorCode: 1,
			Data:      nil,
		})
		return
	}
	if user.Password != Services.GenerateHash(login.Password) {
		c.JSON(http.StatusBadRequest, Models.Response{
			Message:   "username or password is incorrect",
			ErrorCode: 1,
			Data:      nil,
		})
		return
	}

	validToken, err := Middleware.GenerateToken(login.Username)
	c.JSON(http.StatusBadRequest, Models.Response{
		Message:   "",
		ErrorCode: 0,
		Data:      validToken,
	})
	fmt.Println(validToken)
}
