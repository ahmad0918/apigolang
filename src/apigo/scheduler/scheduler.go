package scheduler

import (
	"apigolang/src/apigo/middleware"
	"apigolang/src/apigo/utils"
	"github.com/carlescere/scheduler"
	"github.com/robfig/cron/v3"
	logger "github.com/sirupsen/logrus"
	"time"
)

func StartScheduler() {
	checkingCache()
	worker := cron.New()

	// Call logging and create new file for daily logging at midnight; Every 1 Day at 00 O'Clock
	_, _ = worker.AddFunc("@midnight", loggingActivity)
}

func loggingActivity() {
	mw := middleware.Middleware{}
	// Call Logging; Every 1 Day at 00 O'Clock
	_, err := scheduler.Every().Day().At("00:00").Run(mw.LoggingActivity)
	if err != nil {
		logger.Printf("Time Logging Activity %v: ", time.Now())
		logger.Printf("Error in Scheduler Logging Activity %v: ", err)
	}
}

func checkingCache() {
	// Check  if already reached its limit
	cache := "SCHEDULER_GOLANG"
	// Check request tries
	totalTries := 1
	var _, err = utils.CheckCacheKeyWithCounter(cache, time.Hour*24, totalTries)
	if err != nil {
		utils.LogError(err, "Error in Insert API Golang")
		return
	}
	defer utils.DeleteCacheKey(cache)
}
