package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaoao98/apiFromScratch/src/controller/api"
)

var CommonRouter []func(router *gin.RouterGroup)

func init() {
	CommonRouter = append(CommonRouter, func(router *gin.RouterGroup) {
		// 测试用例1
		router.GET("ping", api.Ping)
		//测试用例2
		router.GET("showMyself", api.ShowMyself)
		//测试用例3
		router.POST("printMyself", api.PrintMyself)
	})
}