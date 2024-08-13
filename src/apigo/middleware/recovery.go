package middleware

import (
	"apigolang/src/apigo/utils"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {

		defer func() {
			err := recover()

			// check if error panic is exist
			if err != nil {
				if _, ok := err.(error); ok {
					stackStr := string(debug.Stack())
					//Write log
					logger.WithFields(logger.Fields{
						"URL":    c.Request.URL,
						"Method": c.Request.Method,
						"Error":  stackStr,
					}).Error("Failed make request")
					utils.Loge.Info("Req to ", c.Request.URL, " | Method=", c.Request.Method, " | Error =>", stackStr)
				} else {
					//Write log
					logger.WithFields(logger.Fields{
						"URL":    c.Request.URL,
						"Method": c.Request.Method,
						"Error":  err,
					}).Error("Failed make request")
					utils.Loge.Info("Req to ", c.Request.URL, " | Method=", c.Request.Method, " | Error =>", err)
				}
				//Write message and abort the request
				c.JSON(http.StatusInternalServerError, gin.H{
					"Message":  "Mohon Maaf, Terjadi kesalahan pada sistem, silakan contact admin ",
					"Response": http.StatusInternalServerError,
					"Result":   nil,
				})
				c.Abort()
			}
		}()

		// execute request
		c.Next()
	}
}
