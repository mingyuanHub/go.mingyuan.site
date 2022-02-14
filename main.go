package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"mingyuanHub/mingyuan.site/utils/logger"
	"mingyuanHub/mingyuan.site/web/controllers"
)

func main() {

	err := logger.InitLogger()
	if err != nil {
		log.Fatalf("fail to initLogger, err=%s", err.Error())
	}

	r := gin.Default()

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