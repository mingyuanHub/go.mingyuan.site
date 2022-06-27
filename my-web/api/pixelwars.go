package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mingyuanHub/mingyuan.site/internal/models"
	"mingyuanHub/mingyuan.site/pkg/template"
	"net/http"
)

func PixelWarsIndex(c *gin.Context) {

	var tmplName = "pixelwars.html"

	t, _ := template.ParseFiles(fmt.Sprintf("my-web/views/%s", tmplName))

	t.ExecuteTemplate(c.Writer, tmplName, nil)
}

func PixelWarsGetColor(c *gin.Context) {
	response := models.NewPixelWarsModel().GetColor()
	c.JSONP(http.StatusCreated, response)
}


func PixelWarsSetColor(c *gin.Context) {

	type Request struct {
		Id int
		Color string
	}

	var (
		req *Request
		err error
	)

	c.ShouldBind(&req)

	pixel := &models.PixelWars{
		Id:      req.Id,
		Color: req.Color,
	}

	pw := models.NewPixelWarsModel()
	err = pw.Save(pixel)

	fmt.Println(pixel, err)

	response := struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}{
		200,
		"ok",
	}
	c.JSONP(http.StatusCreated, response)
}