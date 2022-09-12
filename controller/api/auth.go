package api

import (
	"GoLang/config"
	"GoLang/models"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	data1 := []byte("hello")
	hash1 := sha256.Sum256(data1)

	data2 := []byte("hello")
	hash2 := sha256.Sum256(data2)

	fmt.Println(hash1 == hash2)
	c.JSON(http.StatusOK, map[string]string{
		"message": "ok",
	})
}

func Register(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Message:   err,
			ErrorCode: 1,
			Data:      nil,
		})
		return
	}

	if user.Password != "" {
		hash2 := sha256.Sum256([]byte(user.Password))
		user.Password = hex.EncodeToString(hash2[:])
	}

	err = config.Database.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusBadGateway, models.Response{
			Message:   err,
			ErrorCode: 2,
			Data:      nil,
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Message:   nil,
		ErrorCode: nil,
		Data:      user,
	})

}
