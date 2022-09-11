package api

import (
	"crypto/sha256"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	//var sampleSecretKey = []byte("SecretYouShouldHide")
	//token := jwt.New(jwt.SigningMethodEdDSA)
	//claims := token.Claims.(jwt.MapClaims)
	//claims["exp"] = time.Now().Add(10 * time.Minute)
	//claims["authorized"] = true
	//claims["user"] = "username"

	//tokenString, err := token.SignedString(sampleSecretKey)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(tokenString)
	//sh1 := sha256.New()
	//sh1.Write([]byte("amin"))
	//
	//sh2 := sha256.New()
	//sh2.Write([]byte("amin"))

	//fmt.Println(sh2.s)
	data1 := []byte("hello")
	hash1 := sha256.Sum256(data1)

	data2 := []byte("hello")
	hash2 := sha256.Sum256(data2)

	fmt.Println(hash1 == hash2)

	c.JSON(http.StatusOK, map[string]string{
		"message": "ok",
	})
}
