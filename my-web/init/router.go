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

	g0 := r.Group("/")
	{
		g0.GET("", api.TimestampIndex)
	}

	g1 := r.Group("/timestamp")
	{
		g1.GET("", api.TimestampIndex)
	}

	g2 := r.Group("/json")
	{
		g2.GET("", api.JsonIndex)
	}

	g3 := r.Group("/bulletscreen")
	{
		g3.GET("", api.BulletScreenIndex)
		g3.GET("/message", api.BulletScreenGetMessage)
		g3.POST("/message", api.BulletScreenSaveMessage)
	}

	g4 := r.Group("/pixelwars")
	{
		g4.GET("", api.PixelWarsIndex)
		g4.GET("/websocket", api.PixelWarsWebsocket)
		g4.GET("/color", api.PixelWarsGetColor)
		g4.POST("/color", api.PixelWarsSetColor)
	}
}