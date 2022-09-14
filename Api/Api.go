package Api

import "github.com/gin-gonic/gin"

var Api *gin.Engine
var Router *gin.RouterGroup

func init() {
	Api = gin.New()
	Router = Api.Group("/Api")
}
