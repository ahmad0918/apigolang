package middleware

import (
	"apigolang/src/apigo/utils"
	"bytes"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

type Request struct{}

// RequestLoggerActivity Will log request method, accessed url, and potentially sent parameter
func (mw *Middleware) RequestLoggerActivity() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "POST" || c.Request.Method == "PUT" {
			buf, _ := ioutil.ReadAll(c.Request.Body)
			rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
			rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf)) //We have to create a new Buffer, because rdr1 will be read.

			re := regexp.MustCompile(`\r?\n`)
			var request = re.ReplaceAllString(readBody(rdr1), "")
			utils.LogActivity(
				c.Request.Method,
				c.Request.URL,
				request,
				"HTTP Request Method",
				requestid.Get(c))
			c.Request.Body = rdr2
			c.Next()
		} else {
			utils.LogActivity(
				c.Request.Method,
				c.Request.URL,
				"",
				"HTTP Request Method",
				requestid.Get(c))
			c.Next()
		}
	}
}

func (mw *Middleware) LoggingActivity() {
	dt := time.Now()
	date := dt.Format("20060102")
	//remember, because we get the path from the CONFIG directory, path.log should cd to the src first
	var filename = viper.GetString("path.log") + "log" + date + ".log"

	// Create the log file if it doesn't exist. And append to it if it already exists.
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777) // src/log/gigapixel/[log + date].log
	if err != nil {
		logger.Fatal(err)
	}
	Formatter := new(logger.TextFormatter)
	// You can change the Timestamp format. But you have to use the same date and time.
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true

	logger.SetFormatter(Formatter)
	logger.SetOutput(f)
}

// Read received request body
func readBody(reader io.Reader) string {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(reader)
	if err != nil {
		return ""
	}

	s := buf.String()
	return s
}
