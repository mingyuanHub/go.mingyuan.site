package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mingyuanHub/mingyuan.site/internal/models"
	"mingyuanHub/mingyuan.site/pkg/template"
	"net/http"
	"strconv"
)

func BulletScreenIndex(c *gin.Context) {

	var tmplName = "bulletscreen.html"

	t, _ := template.ParseFiles(fmt.Sprintf("my-web/views/%s", tmplName))

	t.ExecuteTemplate(c.Writer, tmplName, nil)
}

func BulletScreenGetMessage(c *gin.Context) {
	offset, _ := strconv.Atoi(c.Query("offset"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	fmt.Println(offset, limit)
	response := models.NewBulletScreenModel().GetBulletScreen(offset, limit)
	c.JSONP(http.StatusCreated, response)
}

func BulletScreenSaveMessage(c *gin.Context) {

	type Request struct {
		Message string
	}

	var (
		req *Request
		err error
	)

	c.ShouldBind(&req)

	bulletScreen := &models.BulletScreen{
		Ip:      c.ClientIP(),
		Message: req.Message,
	}

	bulletScreenModel := models.NewBulletScreenModel()
	err = bulletScreenModel.Save(bulletScreen)

	fmt.Println(bulletScreen, err)

	response := struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}{
		200,
		"ok",
	}
	c.JSONP(http.StatusCreated, response)
}
