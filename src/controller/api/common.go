package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReqShowMyself struct{
	ParaA         string `form:"para_a" binding:"omitempty"`   
	ParaB         string `form:"para_b" binding:"omitempty"`   
}

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func ShowMyself(ctx *gin.Context) {
	var req ReqShowMyself
	if err := ctx.BindQuery(&req); err != nil {
        // 如果绑定过程中发生错误，返回400状态码和错误信息
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	ctx.JSON(http.StatusOK, gin.H{
		"retA": req.ParaA,
		"retB": req.ParaB,
	})
}