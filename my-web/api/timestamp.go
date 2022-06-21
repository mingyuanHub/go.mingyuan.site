package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mingyuanHub/mingyuan.site/pkg/template"
)

func TimestampIndex(c *gin.Context) {

	var tmplName = "timestamp.tmpl"

	t, _ := template.ParseFiles(fmt.Sprintf("my-web/views/%s", tmplName))

	data := map[string]string{
		"now":"home",
	}

	t.ExecuteTemplate(c.Writer, tmplName, data)
}