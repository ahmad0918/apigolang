package utils

import (
	"fmt"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"time"
)

func LoggingActivity() *logger.Logger {
	timNow := time.Now()
	date := timNow.Format("02012006")

	fileName := viper.GetString("path.log") + date + ".log"

	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		LogError(fmt.Errorf("failed to openFile logs: %v", err), err.Error())
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
