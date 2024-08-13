package middleware

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)

		defer func() {
			if errors.Is(context.DeadlineExceeded, ctx.Err()) && !c.Writer.Written() {

				c.JSON(http.StatusGatewayTimeout, gin.H{
					"status":  http.StatusGatewayTimeout,
					"message": "Server Timeout",
				})
				c.Abort()
			}

			cancel()
		}()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
