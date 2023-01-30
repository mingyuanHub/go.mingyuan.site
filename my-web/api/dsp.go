package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mingyuanHub/mingyuan.site/internal/models"
	"mingyuanHub/mingyuan.site/pkg/openrtb"
	"mingyuanHub/mingyuan.site/pkg/setting"
	"mingyuanHub/mingyuan.site/pkg/template"
	"net/http"
	"strconv"
	"time"
)

func AdxIndex(c *gin.Context) {

	var tmplName = "adx.html"

	t, _ := template.ParseFiles(fmt.Sprintf("my-web/views/%s", tmplName))

	t.ExecuteTemplate(c.Writer, tmplName, nil)
}

func AdxDspSave(c *gin.Context) {

	var (
		dsp *models.Dsp
		err error
	)

	c.ShouldBind(&dsp)

	pw := models.NewDspModel()
	err = pw.Save(dsp)

	fmt.Println(dsp, err)

	response := struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}{
		200,
		"ok",
	}
	c.JSONP(http.StatusOK, response)
}

func AdxGetDspList(c *gin.Context) {
	dspList := models.NewDspModel().GetDspList()

	for _, dsp := range dspList {
		dsp.Adm = ""
	}

	c.JSONP(http.StatusOK, dspList)
}

func AdxGetDspAdm(c *gin.Context) {

	id, _ := strconv.Atoi(c.Query("id"))

	fmt.Println(id)

	adm := models.NewDspModel().GetDspAdmById(id)

	c.JSONP(http.StatusOK, adm)
}

func AdxSaveDspNotice(c *gin.Context) {
	uniqueKey := c.Params.ByName("uniqueKey")
	noticeType := c.Params.ByName("noticeType")

	dsp := models.NewDspModel().GetDspByUniqueKey(uniqueKey)

	noticeTypeInt, err := models.NewDspNoticeModel().GetNoticeTypeValue(noticeType)

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

	fmt.Println(uniqueKey, noticeType, dsp)

	netIp, _ := c.RemoteIP()
	ua := c.Request.Header.Get("User-Agent")

	dspNotice := &models.DspNotice{
		DspId:      dsp.Id,
		NoticeType: noticeTypeInt,
		Ip:         netIp.String(),
		Ua:         ua,
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
	}

	err = models.NewDspNoticeModel().Save(dspNotice)

	if err != nil {
		fmt.Println("fail to save dspNotice, err=", err.Error())
	}

	response := struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}{
		200,
		"ok",
	}

	c.JSONP(http.StatusOK, response)
}

func AdxGetDspNotice(c *gin.Context) {

	dspId, _ := strconv.Atoi(c.Query("id"))
	noticeId, _ := strconv.Atoi(c.Query("noticeId"))
	nowDate := c.Query("nowDate")

	dspNoticeList := models.NewDspNoticeModel().GetDspNoticeByDspId(dspId, noticeId, nowDate)
	
	type dataIterm struct {
		*models.DspNotice
		
		NoticeTypeName string `json:"notice_type_name"`
	}

	var data = []*dataIterm{}
	for _, dspNotice := range dspNoticeList {

		//dspNotice.CreateTime =

		var dtIterm = &dataIterm{
			DspNotice:dspNotice,
		}

		noticeTypeName, err := models.NewDspNoticeModel().GetNoticeTypeName(dspNotice.NoticeType)
		if err != nil {
			fmt.Println(err)
			continue
		}

		dtIterm.NoticeTypeName = noticeTypeName

		data = append(data, dtIterm)
	}

	c.JSONP(http.StatusOK, data)
}

func AdxBid(c *gin.Context) {
	var adRequest = &openrtb.AdRequest{}

	c.ShouldBind(&adRequest)

	requestId := adRequest.Id

	uniqueKey := c.Params.ByName("uniqueKey")

	dsp := models.NewDspModel().GetDspByUniqueKey(uniqueKey)

	adResponse := openrtb.NewAdResponse()

	adResponse.Id = requestId
	adResponse.SeatBid[0].Bid[0].Id = adRequest.Id
	adResponse.SeatBid[0].Bid[0].Price = 99.99
	adResponse.SeatBid[0].Bid[0].NUrl = setting.AppConfig.LocalHost + "adx/" + uniqueKey + "/nurl"
	adResponse.SeatBid[0].Bid[0].BUrl = setting.AppConfig.LocalHost + "adx/" + uniqueKey + "/burl"
	adResponse.SeatBid[0].Bid[0].LUrl = setting.AppConfig.LocalHost + "adx/" + uniqueKey + "/lurl"
	adResponse.SeatBid[0].Bid[0].Adm = dsp.Adm
	adResponse.SeatBid[0].Bid[0].Ext.NUrl = []string{setting.AppConfig.Host + "adx/" + uniqueKey + "/tpnurl"}
	adResponse.SeatBid[0].Bid[0].Ext.LUrl = []string{setting.AppConfig.Host + "adx/" + uniqueKey + "/tplurl"}
	adResponse.SeatBid[0].Bid[0].Ext.BUrl = []string{setting.AppConfig.Host + "adx/" + uniqueKey + "/tpburl"}
	adResponse.SeatBid[0].Bid[0].Ext.ImpUrl = []string{setting.AppConfig.Host + "adx/" + uniqueKey + "/tpimpurl"}
	adResponse.SeatBid[0].Bid[0].Ext.ClkUrl = []string{setting.AppConfig.Host + "adx/" + uniqueKey + "/tpclkurl"}

	c.JSONP(http.StatusOK, adResponse)
}


