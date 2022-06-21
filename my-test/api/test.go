package api

import (
	"github.com/gin-gonic/gin"
	"mingyuanHub/mingyuan.site/pkg/logger"
	"net/http"
)

func Test1(c *gin.Context) {

	defer func() {
		if panicMsg := recover(); panicMsg != nil {
			logger.Error("[home] error: recover from Home. msgError=%v", panicMsg)
		}
	}()

	//t, err := template.ParseFiles("my-web/views/_header.tmpl", "my-web/views/_footer.tmpl", "my-web/views/home.tmpl")
	//
	//if err != nil {
	//	logger.Error("[home] error: %s", err.Error())
	//	return
	//}

	//err = t.ExecuteTemplate(c.Writer, "home.tmpl", data)
	//
	//if err != nil {
	//	logger.Error("[home] error: %s", err.Error())
	//}

	//c.ExecuteTemplate(os.Stdout, "home.tmpl", data)

	data := map[string]string{
		"content":"home",
	}

	c.HTML(http.StatusOK, "home.tmpl", data)
}

func Test2(c *gin.Context) {

	data := map[string]string{
		"content":"exchange",
	}

	c.JSONP(http.StatusCreated, data)
}