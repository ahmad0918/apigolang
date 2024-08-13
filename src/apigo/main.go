package main

import (
	"apigolang/src/apigo/connection"
	"apigolang/src/apigo/scheduler"
	"apigolang/src/apigo/utils"
	"apigolang/src/config/routers"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// @title APIGOLANG
// @version 1.0
// @description This is a sample server for Apigolang application.

// @contact.name Ahmad Hilmy Muflih
// @contact.email ahmadhilmy0918123@gmail.com
// @schemes http
// @host localhost:8080
// @BasePath /api/golang/public
func init() {
	viper.SetConfigFile("src/config/config.json")
	err := viper.ReadInConfig()
	if err != nil {
		utils.LogError(err, "failed to read config: %v")
	}
}

func main() {
	//close all pool connection
	defer connection.InitCloseConnection()

	router := gin.New()
	routers.NewRouters(router)
	utils.LoggingActivity()

	host := viper.GetString("host")
	addr := ":" + viper.GetString("port")

	c := cors.AllowAll()
	handler := c.Handler(router)
	server := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 90 * time.Second,
	}

	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		scheduler.StartScheduler()
		<-quit
		utils.LogSuccess(nil, "Server Apigolang is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			utils.LogError(nil, fmt.Sprintf("Could not gracefully shutdown the server Apigolang: %v\n", err))
		}
		close(done)
	}()

	utils.LogSuccess(nil, fmt.Sprintf("Server Apigolang is ready to handle requests at http://%s%s", host, addr))
	if errS := server.ListenAndServe(); errS != nil && !errors.Is(http.ErrServerClosed, errS) {
		logger.Fatalf("Could not listen on %s: %v\n", host, errS)
	}

	<-done
	utils.LogSuccess(nil, "Server Apigolang stopped")

}
