package controllers

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"mingyuanHub/mingyuan.site/utils/logger"
)

func Home(c *gin.Context) {

	logger.Info("home api request: %s", c.Request.UserAgent())

	t, err := template.ParseFiles("web/views/_header.tmpl", "web/views/_footer.tmpl", "web/views/home.tmpl")

	if err != nil {
		logger.Info("home error: %s", err.Error())

		c.JSON(200, nil)
		return
	}

	data := map[string]string{
		"content":"home",
	}

	err = t.ExecuteTemplate(c.Writer, "home.tmpl", data)

	if err != nil {
		logger.Info("home error: %s", err.Error())

		c.JSON(200, nil)
		return
	}

	//t.ExecuteTemplate(os.Stdout, "home.tmpl", data)

	//c.HTML(http.StatusOK, "home.tmpl", gin.H{
	//	"title": "home website",
	//})
}