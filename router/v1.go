package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yueqingkong/IP-Proxy-Go/conf"
	"github.com/yueqingkong/IP-Proxy-Go/router/middleware"
	"log"
)

func HttpServer() {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(middleware.ErrorHandler())

	v1(r)

	port := fmt.Sprintf(":%s", conf.Port())
	err := r.Run(port)
	if err != nil {
		log.Fatal(err)
	}
}

func v1(engine *gin.Engine) {
	v1 := engine.Group("/proxy/v1")

	brows := v1.Group("/brows")
	brows.GET("/ip", IP)
}
