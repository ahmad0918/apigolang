package middleware

import (
	"github.com/spf13/viper"
	"github.com/unrolled/secure"
)

// UnrolledSecure is a function to secure process
func UnrolledSecure() *secure.Secure {
	var isDev bool
	if env := viper.GetString("environment"); env == "development" {
		isDev = true
	} else {
		isDev = false
	}

	return secure.New(secure.Options{
		FrameDeny:               true,
		ContentSecurityPolicy:   "frame-ancestors 'none';",
		ReferrerPolicy:          "strict-origin-when-cross-origin",
		CrossOriginOpenerPolicy: "same-origin",
		ContentTypeNosniff:      true,
		IsDevelopment:           isDev,
	})
}
