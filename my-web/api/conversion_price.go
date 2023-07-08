package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mingyuanHub/mingyuan.site/pkg/encrypt"
	"mingyuanHub/mingyuan.site/pkg/template"
	"net/http"
)

func ConversionPrice(c *gin.Context) {
	var tmplName = "conversion_price.html"
	t, _ := template.ParseFiles(fmt.Sprintf("my-web/views/%s", tmplName))
	t.ExecuteTemplate(c.Writer, tmplName, nil)
}

func PriceEncrypt(c *gin.Context) {
	type Request struct {
		Type int
		Price string
		Secret string
	}

	var (
		req *Request
		err error
	)

	c.ShouldBind(&req)

	fmt.Println(encrypt.VerifyEncryptKey(req.Secret))

	if b, _ := encrypt.VerifyEncryptKey(req.Secret); !b {
		response := struct {
			Status  int    `json:"status"`
			Message string `json:"message"`
		}{
			400,
			"invalid secret",
		}
		c.JSONP(http.StatusOK, response)
		return
	}


	var (
		priceStr string
	)

	if req.Type == 1 {
		priceStr = encrypt.EcpmEncrypt(req.Price, req.Secret)
	} else {
		priceStr, err = encrypt.EcpmDecrypt(req.Price, req.Secret)
		if err != nil {
			response := struct {
				Status  int    `json:"status"`
				Message string `json:"message"`
			}{
				401,
				err.Error(),
			}
			c.JSONP(http.StatusOK, response)
			return
		}
	}

	res := map[string]interface{}{
		"status": 200,
		"data":   priceStr,
	}

	c.JSONP(http.StatusOK, res)
}
