package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mingyuanHub/mingyuan.site/pkg/template"
)

func JsonIndex(c *gin.Context) {

	var tmplName = "json.tmpl"

	t, _ := template.ParseFiles(fmt.Sprintf("my-web/views/%s", tmplName))

	data := map[string]string{
		"now":"home",
	}

	t.ExecuteTemplate(c.Writer, tmplName, data)
}