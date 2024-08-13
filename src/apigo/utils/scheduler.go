package utils

import (
	"fmt"
	"time"
)

func CheckCacheKeyWithCounter(key string, wait time.Duration, times int) (int, error) {
	var tries int
	exist, ok := mainCache.Get(key)
	if ok {
		tries = exist.(int)
		LogSuccess(tries, "TotalTries: ")
		if tries >= times {
			return 0, fmt.Errorf("user tries more than %d times", times)
		} else {
			inc, _ := mainCache.IncrementInt(key, 1)
			return inc, nil
		}
	} else {
		tries = 1
		mainCache.Set(key, tries, wait)
	}
	return tries, nil
}

func DeleteCacheKey(key string) {
	_, expTm, ok := mainCache.GetWithExpiration(key)
	if !ok {
		mainCache.Delete(key)
		LogSuccess("ExpiredTime: "+expTm.Format("2006-01-02 15:04:05"), "Key "+key+" Expired")
	}
	LogSuccess("ExpiredTime: "+expTm.Format("2006-01-02 15:04:05"), "Key "+key+" Still Active")
}
