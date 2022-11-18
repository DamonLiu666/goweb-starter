package router

import (
	"github.com/gin-gonic/gin"
	"yy-data-processing/apis"
)

func init() {
	routersNoCheck = append(routersNoCheck, registerUserRouter)
}

func registerUserRouter(v1 *gin.RouterGroup) {
	apis := apis.UserApis{}
	r := v1.Group("/user")
	{
		r.POST("/createUser", apis.SaveUser)
		r.GET("/findUser", apis.FindUser)
	}
}
