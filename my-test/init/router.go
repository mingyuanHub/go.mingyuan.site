package init

import (
	"github.com/gin-gonic/gin"
	"mingyuanHub/mingyuan.site/internal/middlewares"
	"mingyuanHub/mingyuan.site/my-test/api"
)

func Init(r *gin.Engine)  {
	initRouter(r)
}

func initRouter(r *gin.Engine)  {
	r.Use(middlewares.LoggerMiddleware())

	r.LoadHTMLGlob("./my-web/views/*")

	r.GET("/home", api.AddArticle)

	g1 := r.Group("/article")
	{
		g1.GET("", api.GetArticles)
		g1.GET("/:id", api.GetArticle)
		g1.POST("", api.AddArticle)
		g1.PUT(":id", api.UpdateArticle)
		g1.DELETE(":id", api.DeleteArticle)
	}
}