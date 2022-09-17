package Middleware

import (
	"GoLang/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

var secretkey string = "i{ktShAS6345cT>1?jgIWxL>BI[4vofTs+zn{5up#Kfo-cDIZ+I#$%^&*(1234s21/aL>ekk*EMa2s3f*"

func IsAuthorized(c *gin.Context) error {
	if c.Request.Header["Token"] == nil {
		c.JSON(http.StatusUnauthorized, Models.Response{
			Message:   "token not found",
			ErrorCode: 10,
		})
		return fmt.Errorf("token not found")
	}

	var mySigningKey = []byte(secretkey)

	token, err := jwt.Parse(c.Request.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing token")
		}
		return mySigningKey, nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, Models.Response{
			Message:   "your token has been expired",
			ErrorCode: 10,
		})
		return fmt.Errorf("your token has been expired")
	}
	if !token.Valid {
		c.JSON(http.StatusUnauthorized, Models.Response{
			Message:   "token not valid",
			ErrorCode: 10,
		})
		return fmt.Errorf("token not valid")
	}
	return nil
}

func GenerateToken(username string) (string, error) {
	var mySigningKey = []byte(secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 240).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return tokenString, err
}
