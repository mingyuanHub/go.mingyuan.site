package base

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type myResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseOk(c *gin.Context, data interface{}) {
	c.JSONP(http.StatusOK, myResponse{
		Status: 200,
		Data:   data,
	})
}

func ResponseError(c *gin.Context, errorMsg string) {
	c.JSONP(http.StatusOK, myResponse{
		Status:  400,
		Message: errorMsg,
	})
}
