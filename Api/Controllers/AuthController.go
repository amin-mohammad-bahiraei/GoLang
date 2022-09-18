package Controllers

import (
	"GoLang/Data"
	"GoLang/Middleware"
	"GoLang/Models"
	"GoLang/Redis"
	"GoLang/Services"
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
)

var db *gorm.DB
var rdb *redis.Client
var ctx context.Context

func init() {
	db = Data.Database
	rdb = Redis.Redis
	ctx = context.Background()
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
}

func GetUserByUsername(c *gin.Context) {
	if err := Middleware.IsAuthorized(c); err != nil {
		return
	}
	username := c.Params.ByName("username")
	val, err := rdb.Get(ctx, username).Result()

	if err != nil {
		var user Models.User
		db.Where("username = ?", username).First(&user)

		strUser, _ := json.Marshal(user)
		rdb.Set(ctx, user.Username, string(strUser), 0)

		c.JSON(http.StatusOK, Models.Response{
			Data: user,
		})
		return
	} else {
		var user Models.User
		json.Unmarshal([]byte(val), &user)
		c.JSON(http.StatusOK, Models.Response{
			Data: user,
		})
		return
	}

}
