package Middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var secretkey string = "this is secret key"

func IsAuthorized(val string) (bool, string) {

	if val == "" {
		return false, "No Token Found"
	}

	var mySigningKey = []byte(secretkey)

	token, err := jwt.Parse(val, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing token.")
		}
		return mySigningKey, nil
	})

	if err != nil {
		fmt.Println("Your Token has been expired.")
		return false, "Your Token has been expired."
	}
	if token.Valid {
		return true, "Token Valid"
	} else {
		return false, "Token Not Valid"
	}
}

func GenerateToken(username string) (string, error) {
	var mySigningKey = []byte(secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return tokenString, err
}
