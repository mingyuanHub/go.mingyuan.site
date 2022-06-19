package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mingyuanHub/mingyuan.site/pkg/template"
	"mingyuanHub/mingyuan.site/pkg/logger"
	"net/http"
)

func Index(c *gin.Context) {

	var tmplName = "timestamp.tmpl"

	t, err := template.ParseFiles(fmt.Sprintf("my-web/views/%s", tmplName))

	if err != nil {
		logger.Error("error: %s", err.Error())
		return
	}

	data := map[string]string{
		"now":"home",
	}

	t.Delims("[[", "]]")

	t.ExecuteTemplate(c.Writer, tmplName, data)
}

func Exchange(c *gin.Context) {

	data := map[string]string{
		"content":"exchange",
	}

	c.JSONP(http.StatusCreated, data)
}