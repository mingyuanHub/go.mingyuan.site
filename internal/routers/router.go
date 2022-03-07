package routers

import (
	"github.com/gin-gonic/gin"
	"mingyuanHub/mingyuan.site/internal/middlewares"
	"mingyuanHub/mingyuan.site/internal/routers/api"
	"mingyuanHub/mingyuan.site/internal/routers/admin"
)

func Init() *gin.Engine {
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

	g1 := r.Group("/article")
	{
		g1.GET("", admin.GetArticles)
		g1.GET("/:id", admin.GetArticle)
		g1.POST("", admin.AddArticle)
		g1.PUT(":id", admin.UpdateArticle)
		g1.DELETE(":id", admin.DeleteArticle)
	}

	return r
}
