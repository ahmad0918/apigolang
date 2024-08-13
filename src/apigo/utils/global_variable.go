package utils

import (
	"time"
	"zgo.at/zcache"

	"github.com/alexcesaro/log/stdlog"
)

var (
	mainCache = zcache.New(time.Minute*10, time.Hour*24)
	Loge      = stdlog.GetFromFlags()
)
