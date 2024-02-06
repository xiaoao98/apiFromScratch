package router

import "github.com/gin-gonic/gin"

var Router *gin.Engine

func init() {
	Router = gin.New()
	Router.Use(gin.Recovery())
	commonGroup := Router.Group("/common")
	for _, f := range CommonRouter {
		f(commonGroup)
	}
}