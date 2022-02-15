package main

import (
	"log"
	"mingyuanHub/mingyuan.site/web/middlewares"

	"github.com/gin-gonic/gin"

	"mingyuanHub/mingyuan.site/utils/logger"
	"mingyuanHub/mingyuan.site/web/controllers"
)

func main() {

	err := logger.Init()
	if err != nil {
		log.Fatalf("fail to initLogger, err=%s", err.Error())
	}

	r := gin.Default()

	r.Use(middlewares.LoggerMiddleware())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.LoadHTMLGlob("web/views/*")

	r.GET("/demo", controllers.Demo)

	r.GET("/home", controllers.Home)

	logger.Info("listen and serve on 0.0.0.0:9080")

	r.Run(":9080")
}