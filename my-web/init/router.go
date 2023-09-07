package init

import (
	"github.com/gin-gonic/gin"
	"mingyuanHub/mingyuan.site/internal/middlewares"
	"mingyuanHub/mingyuan.site/my-web/api"
	"mingyuanHub/mingyuan.site/my-web/api/chatgpt"
	"mingyuanHub/mingyuan.site/my-web/api/ip"
)

func Init(r *gin.Engine)  {
	initRouter(r)
}

func initRouter(r *gin.Engine) {
	r.Use(middlewares.LoggerMiddleware())

	r.LoadHTMLGlob("./my-web/views/*")

	//r.Static("/assets", "./assets")
	//r.StaticFS("/more_static", http.Dir("/var/log"))
	//r.StaticFile("/bdunion.txt", "./assets/bdunion.txt")

	g0 := r.Group("/")
	{
		g0.GET("", api.ConversionTimestamp)
		g0.GET("/timestamp", api.ConversionTimestamp)
		g0.GET("/json", api.ConversionJson)
		g0.GET("/md5", api.ConversionMd5)
		g0.GET("/url", api.ConversionUrl)
		g0.GET("/base64", api.ConversionBase64)
		g0.GET("/utf8", api.ConversionUtf8)
		g0.GET("/trans", api.ConversionTrans)
		g0.GET("/trans/youdao", api.ConversionTransYoudao)
		g0.GET("/markdown", api.MarkDown)
		g0.GET("/timeline", api.Timeline)

		g0.GET("/data-diff", api.DataDiff)
		g0.GET("/data-filter", api.DataFilter)
		g0.GET("/data-calculator", api.DataCalculcator)
		g0.POST("/data-calc", api.DataCalc)

		g0.GET("/json2gostruct", api.DevJson2GoStruct)
		g0.POST("/json2gostruct-api", api.DevJson2GoStructApi)

		g0.GET("/aes", api.ConversionAes)
		g0.GET("/price", api.ConversionPrice)
		g0.POST("/price/price-encrypt", api.PriceEncrypt)
		g0.GET("/gepip", api.GeoIp)

		g0.GET("/map", api.Map)
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

	g5 := r.Group("/adx")
	{
		g5.GET("", api.AdxIndex)

		g5.POST("/adxGetDspList", api.AdxGetDspList)
		g5.GET("/adxGetDspAdm", api.AdxGetDspAdm)

		g5.POST("/adxDspSave", api.AdxDspSave)
		g5.GET("/adxGetDspNotice", api.AdxGetDspNotice)

		g5.GET("/:uniqueKey/:noticeType", api.AdxSaveDspNotice)
		g5.POST("/:uniqueKey", api.AdxBid)
	}

	g6 := r.Group("/chatgpt")
	{
		g6.GET("", chatgpt.Index)
		g6.POST("/checkToken", chatgpt.CheckToken)
	}

	g7 := r.Group("/ip")
	{
		g7.GET("", ip.Index)
		g7.GET("/search", ip.Search)
	}
}