package controllers

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"mingyuanHub/mingyuan.site/utils/logger"
)

func Home(c *gin.Context) {

	defer func() {
		if panicMsg := recover(); panicMsg != nil {
			logger.Error("[home] error: recover from panic in homeController. msgError=%v", panicMsg)
		}
	}()

	t, err := template.ParseFiles("web/views/_header.tmpl", "web/views/_footer.tmpl", "web/views/home.tmpl")

	if err != nil {
		logger.Error("[home] error: %s", err.Error())
		return
	}

	data := map[string]string{
		"content":"home",
	}

	err = t.ExecuteTemplate(c.Writer, "home.tmpl", data)

	if err != nil {
		logger.Error("[home] error: %s", err.Error())
	}

	//t.ExecuteTemplate(os.Stdout, "home.tmpl", data)

	//c.HTML(http.StatusOK, "home.tmpl", gin.H{
	//	"title": "home website",
	//})
}