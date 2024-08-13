package delivery

import (
	_ "apigolang/src/apigo/docs"
	"apigolang/src/apigo/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Routers struct {
	engine *gin.Engine
}

func NewRoutersApiKioskWeb(router *gin.Engine) *Routers {
	return &Routers{
		engine: router,
	}
}

func (r *Routers) Router() {
	//cfgBafKiosk := db.CreateConnection()
	//connBafKiosk, err := connection.NewPostgresConnection(cfgBafKiosk)
	//if err != nil {
	//	utils.LogError(err, err.Error())
	//	return
	//}

	// Connection Param/General Setting
	//newParamSettingRepo := repoParamSetting.NewParamSettingRepo(connBafKiosk.SqlDb(), connSandia.SqlDb())
	//newParamSettingService := serviceParamSetting.NewParamSettingService(newParamSettingRepo)
	//newParamSettingControllers := controllers.NewParamSettingControllers(newParamSettingService)

	mw := middleware.Middleware{}
	r.engine.Use(mw.RequestLoggerActivity())
	urlSwagger := ginSwagger.URL(viper.GetString("host") + "/api/golang/public/v1/web/swagger/doc.json")
	v1Private := r.engine.Group("/api/golang/public/v1/web")
	{
		// Swagger
		v1Private.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, urlSwagger))
	}
}
