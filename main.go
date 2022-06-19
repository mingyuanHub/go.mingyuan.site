package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	webint "mingyuanHub/mingyuan.site/my-web/init"
	"mingyuanHub/mingyuan.site/internal/models"
	"mingyuanHub/mingyuan.site/pkg/logger"
	"mingyuanHub/mingyuan.site/pkg/setting"
)

func main() {
	var err error

	err = logger.Init()
	if err != nil {
		log.Fatalf("fail to initLogger, err=%s", err.Error())
	}

	err = setting.Init()
	if err != nil {
		log.Fatalf("fail to initSetting, err=%s", err.Error())
	}

	err = models.Init()
	if err != nil {
		log.Fatalf("fail to initDB, err=%s", err.Error())
	}

	gin.SetMode(setting.AppConfig.RunMode)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerConfig.ServerPort),
		Handler:        getRouter(),
		ReadTimeout:    time.Duration(setting.ServerConfig.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(setting.ServerConfig.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	logger.Info("listen and serve on 0.0.0.0:%d", setting.ServerConfig.ServerPort)

	//r.Run(":9080")
	if err = s.ListenAndServe(); err != nil {
		log.Fatalf("fail to listenAndServe, err=%s", err.Error())
	}
}

func getRouter() (r *gin.Engine) {
	r = gin.Default()

	webint.Init(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Static("/assets", "./assets")

	return
}
