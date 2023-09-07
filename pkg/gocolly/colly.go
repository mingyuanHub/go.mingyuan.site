package gocolly

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"net"
	"net/http"
	"time"
)

func NewColly() *colly.Collector {
	c := colly.NewCollector(
		//colly.AllowedDomains("www.mfa.gov.cn", "www.fmprc.gov.cn"),
	)

	rule := &colly.LimitRule{
		RandomDelay: time.Second,
		Parallelism: 10, //并发数为10
	}
	_ = c.Limit(rule)

	// 使用随机user-agent
	extensions.RandomUserAgent(c)

	//transport := &http.Transport{
	//	MaxIdleConns:        60000,
	//	MaxConnsPerHost:     10000,
	//	MaxIdleConnsPerHost: 10000,
	//	IdleConnTimeout:     time.Duration(30) * time.Second,
	//	TLSHandshakeTimeout: 5 * time.Second,
	//
	//	DialContext: (&net.Dialer{
	//		Timeout:   30 * time.Second,
	//		KeepAlive: 30 * time.Second,
	//	}).DialContext,
	//}

	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment, // 使用代理
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second, // 超时时间
			KeepAlive: 30 * time.Second, // keepAlive 超时时间
		}).DialContext,
		MaxIdleConns:          100,              // 最大空闲连接数
		IdleConnTimeout:       90 * time.Second, // 空闲连接超时
		TLSHandshakeTimeout:   90 * time.Second, // TLS 握手超时
		ExpectContinueTimeout: 90 * time.Second,
	}

	// HTTP 的配置
	c.WithTransport(transport)

	// 设置请求信息
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("X-Requested-With", "XMLHttpRequest")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3314.0 Safari/537.36 SE 2.X MetaSr 1.0")
	})

	// 请求响应
	c.OnResponse(func(r *colly.Response) {
		//响应状态码, 响应内容
		//fmt.Println(r.StatusCode, string(r.Body))
		fmt.Println(r.Request.URL)
	})

	c.OnError(func(response *colly.Response, err error) {
		fmt.Println("colly OnError: ", err.Error())
	})

	return c
}
