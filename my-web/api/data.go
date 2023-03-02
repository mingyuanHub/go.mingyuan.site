package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mingyuanHub/mingyuan.site/pkg/template"
)

func DataDiff(c *gin.Context) {

	var tmplName = "data_diff.html"

	t, _ := template.ParseFiles(fmt.Sprintf("my-web/views/%s", tmplName))

	t.ExecuteTemplate(c.Writer, tmplName, nil)
}

func DataFilter(c *gin.Context) {

	var tmplName = "data_filter.html"

	t, _ := template.ParseFiles(fmt.Sprintf("my-web/views/%s", tmplName))

	t.ExecuteTemplate(c.Writer, tmplName, nil)
}