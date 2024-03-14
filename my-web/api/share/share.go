package share

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"mingyuanHub/mingyuan.site/internal/models"
	"mingyuanHub/mingyuan.site/my-web/api/base"
	"mingyuanHub/mingyuan.site/pkg/common"
)

func SaveData(c *gin.Context) {
	type request struct {
		Uuid  string      `json:"uuid"`
		Path  string      `json:"path"`
		Name  string      `json:"name"`
		Token string      `json:"token"`
		Data  interface{} `json:"data"`
	}

	var (
		req *request
		err error
	)

	if err = c.ShouldBind(&req); err != nil {
		c.JSONP(http.StatusOK, err.Error())
		return
	}

	share := &models.Share{
		Uuid:  req.Uuid,
		Path:  req.Path,
		Name:  req.Name,
		Token: req.Token,
		Data:  common.String(req.Data),
	}

	shareModel := models.NewShareModel()
	err = shareModel.Save(share)

	if err != nil {
		base.ResponseError(c,err.Error())
		return
	}

	base.ResponseOk(c,nil)
}

func GetData(c *gin.Context) {
	type request struct {
		Uuid  string `json:"uuid"`
		Path  string `json:"path"`
		Token string `json:"token"`
	}

	var (
		req *request
		err error
	)

	if err = c.ShouldBind(&req); err != nil {
		c.JSONP(http.StatusOK, err.Error())
		return
	}

	data := models.NewShareModel().GetShareData(req.Path, req.Token)
	if len(data) == 0 {
		base.ResponseError(c,"expired token")
		return
	}

	base.ResponseOk(c,data)
}