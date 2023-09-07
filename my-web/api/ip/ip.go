package ip

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mingyuanHub/mingyuan.site/pkg/logger"
	"mingyuanHub/mingyuan.site/pkg/net"
	"mingyuanHub/mingyuan.site/pkg/template"
	"net/http"
)

func Index(c *gin.Context) {

	var tmplName = "ip.html"

	t, _ := template.ParseFiles(fmt.Sprintf("my-web/views/%s", tmplName))

	t.ExecuteTemplate(c.Writer, tmplName, nil)
}

func Search(c *gin.Context) {

	ip := c.Query("ip")
	url := fmt.Sprintf("https://qifu-api.baidubce.com/ip/geo/v1/district?ip=%s", ip)
	fmt.Println(url)

	response, body, err := net.HttpGetRequest(url, net.HttpClient5000)
	if err != nil {
		logger.Error("fail to HttpGetRequest, err=%s", err.Error())
		return
	}

	if response.StatusCode != http.StatusOK {
		logger.Error("fail to HttpGetRequest, err=%s", response.StatusCode)
	}

	res := map[string]interface{}{
		"status": 200,
		"data":   string(body),
	}

	c.JSONP(http.StatusOK, res)
}

//newColly := gocolly.NewColly()
//
//var region string
//
//newColly.OnHTML(`tbody`, func(e *colly.HTMLElement) {
//	e.ForEach("span", func(idx int, e *colly.HTMLElement) {
//		if idx == 1 {
//			region = e.Text
//		}
//		fmt.Println(100000, idx, e.Text)
//	})
//})
//newColly.Visit(url)
