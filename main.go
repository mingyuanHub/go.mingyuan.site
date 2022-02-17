package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"mingyuanHub/mingyuan.site/internal/routers"
	"net/http"

	"mingyuanHub/mingyuan.site/internal/pkg/logger"
)

func main() {

	err := logger.Init()
	if err != nil {
		log.Fatalf("fail to initLogger, err=%s", err.Error())
	}

	r := routers.InitRouter()

	logger.Info("listen and serve on 0.0.0.0:9080")

	gin.SetMode(gin.DebugMode)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", 9080),
		Handler:        r,
		//ReadTimeout:    1000,
		//WriteTimeout:   1000,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

	//r.Run(":9080")
}