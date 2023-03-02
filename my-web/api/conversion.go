package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"mingyuanHub/mingyuan.site/pkg/logger"
	"mingyuanHub/mingyuan.site/pkg/net"
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

func ConversionTrans(c *gin.Context) {
	var tmplName = "conversion_trans.html"
	t, _ := template.ParseFiles(fmt.Sprintf("my-web/views/%s", tmplName))
	t.ExecuteTemplate(c.Writer, tmplName, nil)
}

func ConversionAes(c *gin.Context) {
	var tmplName = "conversion_aes.html"
	t, _ := template.ParseFiles(fmt.Sprintf("my-web/views/%s", tmplName))
	t.ExecuteTemplate(c.Writer, tmplName, nil)
}

type TransResponse struct {
	Type        string
	ErrorCode   int
	ElapsedTime int64
	TranslateResult [][]*TranslateResultItermRes
}

type TranslateResultItermRes struct {
	Src string
	Tgt string
}

func ConversionTransYoudao(c *gin.Context) {

	t := c.Query("t")

	if len(t) == 0 {
		return
	}

	var errMsg = errors.New("")

	httpResponse, body, err := net.HttpGetRequest("https://fanyi.youdao.com/translate?&doctype=json&type=AUTO&i="+t, net.HttpClient5000)
	if err != nil {
		logger.Error("fail to HttpGetRequest, err=%s", err.Error())
		return
	}

	if httpResponse.StatusCode != http.StatusOK {
		return
	}

	var transResponse = &TransResponse{}

	//"type":"ZH_CN2EN","errorCode":0,"elapsedTime":1,"translateResult":[[{"src":"12313133","tgt":"12313133"}]]}

	err = json.Unmarshal(body, &transResponse)
	if err != nil {
		errMsg = errors.New(fmt.Sprintf("fail to trans. json unmarshal error='%v'", err))
		return
	}

	var transType, transTgt string

	if transResponse.TranslateResult != nil && len(transResponse.TranslateResult) > 0 {
		if transResponse.TranslateResult[0] != nil && len(transResponse.TranslateResult[0]) > 0 {
			transType = transResponse.Type

			if len(transResponse.TranslateResult[0]) == 1 {
				transTgt = transResponse.TranslateResult[0][0].Tgt
			}

			if len(transResponse.TranslateResult[0]) > 1 {
				for idx, iterm := range transResponse.TranslateResult[0] {
					k := ""
					if idx > 0 {
						k = "\r\n"
					}
					transTgt += fmt.Sprintf("%s%s\r\n%s", k, iterm.Src, iterm.Tgt)
				}
			}
		}
	}

	response := &Response{
		Status:  http.StatusOK,
		Message: errMsg.Error(),
		Result: map[string]interface{}{
			"transType": transType,
			"transTgt":  transTgt,
		},
	}

	c.JSONP(http.StatusOK, response)
}