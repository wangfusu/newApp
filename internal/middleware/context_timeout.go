package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

// ContextTimeout 统一超时管理
func ContextTimeout(t time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
