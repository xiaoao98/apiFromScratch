package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/xiaoao98/apiFromScratch/src/router"
)

var mainCtx, mainCancel = context.WithCancel(context.Background())

// HandleSignal 信号处理
func HandleSignal() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		sig := <-ch
		switch sig {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			// library.ZapLogger.Infof("收到信号，执行平滑退出")
			mainCancel()
		}
	}()
}

func main() {
	HandleSignal()
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "pong",
	// })
	// })
  	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	gin.SetMode(gin.DebugMode)
	go router.Router.Run(":12888")
	<-mainCtx.Done()
}