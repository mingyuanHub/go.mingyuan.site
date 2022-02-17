package routers

import (
	"github.com/gin-gonic/gin"
	"mingyuanHub/mingyuan.site/internal/middlewares"
	"mingyuanHub/mingyuan.site/internal/web/api"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.LoggerMiddleware())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/test", api.Test)

	r.LoadHTMLGlob("./internal/web/views/*")

	r.GET("/home", api.Home)

	return r
}
