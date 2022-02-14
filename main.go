package main

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"mingyuanHub/mingyuan.site/web/controllers"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.LoadHTMLGlob("web/views/*")

	r.GET("/demo", controllers.Demo)

	r.GET("/home", controllers.Home)

	fmt.Println("listen and serve on 0.0.0.0:9080")

	r.Run(":9080")
}