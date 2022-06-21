package init

import (
	"github.com/gin-gonic/gin"
	"mingyuanHub/mingyuan.site/internal/middlewares"
	"mingyuanHub/mingyuan.site/my-web/api"
)

func Init(r *gin.Engine)  {
	initRouter(r)
}

func initRouter(r *gin.Engine) {
	r.Use(middlewares.LoggerMiddleware())

	r.LoadHTMLGlob("./my-web/views/*")

	g1 := r.Group("/timestamp")
	{
		g1.GET("", api.TimestampIndex)
	}

	g2 := r.Group("/json")
	{
		g2.GET("", api.JsonIndex)
	}

}