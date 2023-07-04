package api

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"mingyuanHub/mingyuan.site/pkg/template"
)

func ConversionJson(c *gin.Context) {
	var tmplName = "conversion_json.html"
	t, _ := template.ParseFiles(fmt.Sprintf("my-web/views/%s", tmplName))
	t.ExecuteTemplate(c.Writer, tmplName, nil)
}

func ConversionTimestamp(c *gin.Context) {
	var tmplName = "conversion_timestamp.html"
	t, _ := template.ParseFiles(fmt.Sprintf("my-web/views/%s", tmplName))
	t.ExecuteTemplate(c.Writer, tmplName, nil)
}

func ConversionMd5(c *gin.Context) {
	var tmplName = "conversion_md5.html"
	t, _ := template.ParseFiles(fmt.Sprintf("my-web/views/%s", tmplName))
	t.ExecuteTemplate(c.Writer, tmplName, nil)
}

func ConversionUrl(c *gin.Context) {
	var tmplName = "conversion_url.html"
	t, _ := template.ParseFiles(fmt.Sprintf("my-web/views/%s", tmplName))
	t.ExecuteTemplate(c.Writer, tmplName, nil)
}



func ConversionAes(c *gin.Context) {
	var tmplName = "conversion_aes.html"
	t, _ := template.ParseFiles(fmt.Sprintf("my-web/views/%s", tmplName))
	t.ExecuteTemplate(c.Writer, tmplName, nil)
}


