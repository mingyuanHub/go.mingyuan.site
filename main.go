package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"mingyuanHub/mingyuan.site/internal/pkg/logger"
	"mingyuanHub/mingyuan.site/internal/pkg/setting"
	"mingyuanHub/mingyuan.site/internal/models"
	"mingyuanHub/mingyuan.site/internal/routers"
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

	logger.Info("listen and serve on 0.0.0.0:%d", setting.ServerConfig.ServerPort)

	gin.SetMode(setting.AppConfig.RunMode)

	r := routers.Init()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerConfig.ServerPort),
		Handler:        r,
		ReadTimeout:    time.Duration(setting.ServerConfig.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(setting.ServerConfig.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	//r.Run(":9080")
	if err = s.ListenAndServe(); err != nil {
		log.Fatalf("fail to listenAndServe, err=%s", err.Error())
	}
}