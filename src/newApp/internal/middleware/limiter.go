package middleware

import (
	"NewApp/pkg/app"
	"NewApp/pkg/errcode"
	"NewApp/pkg/limiter"
	"github.com/gin-gonic/gin"
)

// RateLimiter 限流控制
func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
