package routers

import (
	"apigolang/src/apigo/apigolangweb/delivery"
	"apigolang/src/apigo/healthcheck"
	"apigolang/src/apigo/middleware"
	"apigolang/src/apigo/utils"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Router interface {
	Router()
}

func NewRouters(router *gin.Engine) {
	var rm Router
	setGinMode()
	setupDefaultMiddleware(router)

	rm = delivery.NewRoutersApiKioskWeb(router)
	rm.Router()

	router.GET("/api/golang/public/healthcheck/database", healthcheck.GetDBHealthCheck)
	healthcheck.Service(router, "Service Apigolang", healthcheck.DefaultConfig(), []healthcheck.Check{})

}

func setGinMode() {
	if IsProduction() {
		gin.SetMode(gin.ReleaseMode)
		utils.Loge.Info("running in Production env")
	} else {
		gin.SetMode(gin.DebugMode)
		utils.Loge.Info("running in Development env")
	}

}

func IsProduction() bool {
	return viper.GetString("environment") == "production"
}

func setupDefaultMiddleware(router *gin.Engine) {
	mwDefault := middleware.InitMiddleware()
	router.Use(requestid.New())
	router.Use(mwDefault.Recovery())
	router.Use(mwDefault.Security())
}
