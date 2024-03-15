package api

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"mingyuanHub/mingyuan.site/pkg/template"
)

func Home(c *gin.Context) {
	var tmplName = "_home.html"
	t, _ := template.ParseFiles(fmt.Sprintf("my-web/views/%s", tmplName))
	t.ExecuteTemplate(c.Writer, tmplName, nil)
}

func ConversionJson(c *gin.Context) {
	var tmplName = "conversion_json.html"
	t, _ := template.ParseFiles(fmt.Sprintf("my-web/views/%s", tmplName))
	t.ExecuteTemplate(c.Writer, tmplName, nil)
}

func ConversionHtml(c *gin.Context) {
	var tmplName = "conversion_html.html"
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

func ConversionBase64(c *gin.Context) {
	var tmplName = "conversion_base64.html"
	t, _ := template.ParseFiles(fmt.Sprintf("my-web/views/%s", tmplName))
	t.ExecuteTemplate(c.Writer, tmplName, nil)
}

func ConversionUtf8(c *gin.Context) {
	var tmplName = "conversion_utf8.html"
	t, _ := template.ParseFiles(fmt.Sprintf("my-web/views/%s", tmplName))
	t.ExecuteTemplate(c.Writer, tmplName, nil)
}

func ConversionAes(c *gin.Context) {
	var tmplName = "tp_aes.html"
	t, _ := template.ParseFiles(fmt.Sprintf("my-web/views/%s", tmplName))
	t.ExecuteTemplate(c.Writer, tmplName, nil)
}

func MarkDown(c *gin.Context) {
	var tmplName = "markdown.html"
	t, _ := template.ParseFiles(fmt.Sprintf("my-web/views/%s", tmplName))
	t.ExecuteTemplate(c.Writer, tmplName, nil)
}

func Timeline(c *gin.Context) {
	var tmplName = "timeline.html"
	t, _ := template.ParseFiles(fmt.Sprintf("my-web/views/%s", tmplName))
	t.ExecuteTemplate(c.Writer, tmplName, nil)
}

func GeoIp(c *gin.Context) {
	var tmplName = "tp_ip.htm"
	t, _ := template.ParseFiles(fmt.Sprintf("my-web/views/%s", tmplName))
	t.ExecuteTemplate(c.Writer, tmplName, nil)
}

func Map(c *gin.Context) {
	var tmplName = "baidu_map.html"
	t, _ := template.ParseFiles(fmt.Sprintf("my-web/views/%s", tmplName))
	t.ExecuteTemplate(c.Writer, tmplName, nil)
}

func ChartSankey(c *gin.Context) {
	var tmplName = "chart_sankey.html"
	t, _ := template.ParseFiles(fmt.Sprintf("my-web/views/%s", tmplName))
	t.ExecuteTemplate(c.Writer, tmplName, nil)
}

