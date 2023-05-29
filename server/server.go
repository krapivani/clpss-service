package server

import (
	"CLPSS/config"
	"CLPSS/handler"
	_ "encoding/base64"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	_ "github.com/prometheus/client_golang/prometheus/promhttp"
	_ "github.com/sirupsen/logrus"
	_ "net/http/pprof"
	_ "strings"
)

func CreateServer(conf config.ConfigInterface) *gin.Engine {

	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()

	//servicePath := util.GetEnv("HYSTRIX_CONTEXT", "")

	router := gin.New()
	router.Use(
		gin.Recovery(),
	)

	externalV1 := router.Group("/external/v1")
	{
		externalV1.GET("/GetInfo", handler.GetInfo)
		externalV1.GET("/GetUsers/:country", handler.GetUsers)
	}

	{
		router.GET("/", handler.GetInfo)
		router.GET("/health", handler.GetHealth)
		//router.GET(servicePath+"/prometheus", gin.WrapH(promhttp.Handler()))
		//router.GET(servicePath+"/odhystrix.stream", gin.WrapH(hystrixStreamHandler))

	}
	return router
}
