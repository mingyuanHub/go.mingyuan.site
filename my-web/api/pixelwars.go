package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"mingyuanHub/mingyuan.site/internal/models"
	"mingyuanHub/mingyuan.site/pkg/template"
	"net/http"
	"strings"
)

func PixelWarsIndex(c *gin.Context) {

	var tmplName = "pixelwars.html"

	t, _ := template.ParseFiles(fmt.Sprintf("my-web/views/%s", tmplName))

	t.ExecuteTemplate(c.Writer, tmplName, nil)
}

func PixelWarsGetColor(c *gin.Context) {
	response := models.NewPixelWarsModel().GetColor()
	c.JSONP(http.StatusOK, response)
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
		Id:    req.Id,
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
	c.JSONP(http.StatusOK, response)
}


var upGrader = websocket.Upgrader{
	CheckOrigin: func (r *http.Request) bool {
		return true
	},
}

var connList = map[int]*websocket.Conn{}

func PixelWarsWebsocket(c *gin.Context) {
	// 完成和Client HTTP >>> WebSocket的协议升级
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	var num int

	for  {
		num = rand.Intn(10000)
		if _, ok := connList[num]; !ok {
			break
		}
	}

	connList[num] = conn

	defer delete(connList, num)
	defer connList[num].Close()

	ip := c.ClientIP()

	for {
		// 接收客户端message
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		msgArr := strings.Split(string(message), ":")

		pixelIdx := msgArr[0]
		pixelColor := msgArr[1]

		msg := fmt.Sprintf("%s:%s:%s", ip, pixelIdx, pixelColor)

		if len(connList) > 0 {
			for conNum, con := range connList {

				//不需要发给自己客户端message
				if conNum == num {
					continue
				}

				// 向其他用户的客户端发送message
				err = con.WriteMessage(mt, []byte(msg))
				if err != nil {
					log.Println("write:", err)
					continue
				}
			}
		}

	}
}