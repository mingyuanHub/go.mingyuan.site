package controllers

import (
	"github.com/gin-gonic/gin"
	"mingyuanHub/mingyuan.site/utils/logger"
	"net/http"
)

func Home(c *gin.Context) {

	logger.Info("home api request: %s", c.Request.UserAgent())

	c.HTML(http.StatusOK, "home.tmpl", gin.H{
		"title": "home website",
	})
}