package router

import "github.com/gin-gonic/gin"

var routersNoCheck = make([]func(v1 *gin.RouterGroup), 0)

func InitRouters() *gin.Engine {
	r := gin.Default()
	noCheck(r)

	return r
}

func noCheck(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	for _, f := range routersNoCheck {
		f(v1)
	}
}
