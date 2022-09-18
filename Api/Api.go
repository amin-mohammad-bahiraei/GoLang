package Api

import (
	"GoLang/Api/Controllers"
	"github.com/gin-gonic/gin"
)

var Api *gin.Engine

//var Router *gin.RouterGroup

func init() {
	Api = gin.New()
	//Router = Api.Group("/Api")

	Api.POST("/Auth/SignUp", Controllers.SignUp)
	Api.POST("/Auth/SignIn", Controllers.SignIn)
	Api.GET("/Auth/GetUserByUsername/:username", Controllers.GetUserByUsername)

	Api.GET("/Product/GetProductList", Controllers.GetProductList)
}
