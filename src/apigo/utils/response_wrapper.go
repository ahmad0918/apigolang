package utils

import (
	"apigolang/src/apigo/models"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

// HandleSuccess is a function to handle successful operations and send a success response to the front
func HandleSuccess(c *gin.Context, data interface{}, msg, massage string) {
	response := models.SuccessResponse{
		Message:  massage,
		Response: http.StatusOK,
		Result:   data,
	}

	c.JSON(http.StatusOK, response)

	reqID := requestid.Get(c)
	Loge.Info(
		"Resp => method =", c.Request.Method,
		"||ID =", reqID,
		"||url =", c.Request.URL,
		"||msg =", msg,
	)
}

// LogSuccess is a function of success process that only generate log to console
func LogSuccess(detail interface{}, msg string) {
	// Assign detail with parameter above and generate console
	loggingActivity().WithFields(logger.Fields{
		"detail": detail,
	}).Info("HTTP Request Method")
	Loge.Info("msg =", msg, "||detail =", detail)
}

// HandleInternalServerError is a function to handle general internal server errors and send an error response to the front
func HandleInternalServerError(c *gin.Context, detail interface{}, msg, massage string) {
	// Assign struct ResponseWrapper with parameter above
	response := models.InternalErrorResponse{
		Message:  massage,
		Response: http.StatusInternalServerError,
		Result:   nil,
	}

	// Use JSON as a response and send the initialized struct above
	c.JSON(http.StatusInternalServerError, response)

	reqID := requestid.Get(c)
	// Assign detail with parameter above and generate console log
	loggingActivity().WithFields(logger.Fields{
		"detail": detail,
		"ID":     reqID,
	}).Error(msg)
	Loge.Info(
		"Resp => method =", c.Request.Method,
		"||ID =", reqID,
		"||url =", c.Request.URL,
		"||msg =", msg,
	)
}

// LogError is a function of failed process that only generate log to console
func LogError(detail interface{}, msg string) {
	// Assign detail with parameter above and generate console
	loggingActivity().WithFields(logger.Fields{
		"detail": detail,
	}).Error(msg)
	Loge.Info("msg =", msg, "||detail =", detail)
}

// LogActivity is a function to create format log contains method, url, and request body. It's created after service hit.
func LogActivity(method, url, request interface{}, msg, requestID string) {
	// Assign detail with parameter above and generate console
	loggingActivity().WithFields(logger.Fields{
		"method":  method,
		"url":     url,
		"request": request,
		"ID":      requestID,
	}).Info(msg)
	Loge.Info("Req  => method =", method,
		"||ID =", requestID,
		"||url =", url,
		"||param =", request)
}

// LogActivity is a function to create format log contains method, url, and request body. It's created after service hit.
func loggingActivity() *logger.Logger {
	timNow := time.Now()
	date := timNow.Format("02012006")

	fileName := viper.GetString("path.log") + date + ".log"

	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		return nil
	}
	Formatter := new(logger.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true

	loggerApiKiosk := logger.New()
	loggerApiKiosk.SetFormatter(Formatter)
	loggerApiKiosk.SetOutput(f)

	return loggerApiKiosk
}

// HandleUnauthorizedError is a function to handle 401 Unauthorized errors and send an error response to the front
func HandleUnauthorizedError(c *gin.Context, detail interface{}, msg string) {
	// Assign struct ResponseWrapper with parameter above
	response := models.UnauthorizedResponse{
		Message:  "Unauthorized access",
		Response: http.StatusUnauthorized,
		Result:   nil,
	}

	// Use JSON as a response and send the initialized struct above
	c.JSON(http.StatusUnauthorized, response)

	reqID := requestid.Get(c)
	// Assign detail with parameter above and generate console log
	loggingActivity().WithFields(logger.Fields{
		"detail": detail,
		"ID":     reqID,
	}).Error(msg)
	Loge.Info(
		"Resp => method =", c.Request.Method,
		"||ID =", reqID,
		"||url =", c.Request.URL,
		"||msg =", msg,
	)
}

// HandleNotFoundError is a function to handle 404 Not Found errors and send an error response to the front
func HandleNotFoundError(c *gin.Context, detail interface{}, msg, massage string) {
	response := models.DataNotFoundResponse{
		Message:  massage,
		Response: http.StatusNotFound,
		Result:   nil,
	}

	c.JSON(http.StatusNotFound, response)

	reqID := requestid.Get(c)
	loggingActivity().WithFields(logger.Fields{
		"detail": detail,
		"ID":     reqID,
	}).Error(msg)
	Loge.Info(
		"Resp => method =", c.Request.Method,
		"||ID =", reqID,
		"||url =", c.Request.URL,
		"||msg =", msg,
	)
}

// HandleBadRequestError is a function to handle 400 Bad Request errors and send an error response to the front
func HandleBadRequestError(c *gin.Context, data, detail interface{}, msg string) {
	response := models.CommonErrorResponse{
		Message:  "Bad request",
		Response: http.StatusBadRequest,
		Result:   data,
	}

	c.JSON(http.StatusBadRequest, response)

	reqID := requestid.Get(c)
	loggingActivity().WithFields(logger.Fields{
		"detail": detail,
		"ID":     reqID,
	}).Error(msg)
	Loge.Info(
		"Resp => method =", c.Request.Method,
		"||ID =", reqID,
		"||url =", c.Request.URL,
		"||msg =", msg,
	)
}
