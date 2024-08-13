package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

// Middleware represent the data-struct for middleware
type Middleware struct {
	secured *secure.Secure
}

// InitMiddleware initialize the middleware
func InitMiddleware() *Middleware {
	return &Middleware{}
}

func (mw *Middleware) Recovery() gin.HandlerFunc {
	return Recovery()
}

func (mw *Middleware) TimeoutMiddleWare(timeoutDuration time.Duration) gin.HandlerFunc {
	return TimeoutMiddleware(timeoutDuration)
}

func (mw *Middleware) Security() gin.HandlerFunc {
	return func(c *gin.Context) {
		mw.secured = UnrolledSecure()
		err := mw.secured.Process(c.Writer, c.Request)

		if err != nil {
			c.Abort()
			return
		}

		// Set header response
		c.Header("Permitted-Cross-Domain-Policies", "by-content-type")
		c.Header("Clear-Site-Data", "cache")
		c.Header("Cross-Origin-Embedded-Policy", "require-corp")
		c.Header("Cross-Origin-Resource-Policy", "same-origin")

		// Avoid header rewrite if response is a redirection.
		if status := c.Writer.Status(); status > 300 && status < 399 {
			c.Abort()
		}
	}
}
