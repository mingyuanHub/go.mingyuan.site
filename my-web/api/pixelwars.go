package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mingyuanHub/mingyuan.site/pkg/template"
)

func PixelWarsIndex(c *gin.Context) {

	var tmplName = "pixelwars.html"

	t, _ := template.ParseFiles(fmt.Sprintf("my-web/views/%s", tmplName))

	t.ExecuteTemplate(c.Writer, tmplName, nil)
}