package middleware

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// CORS middleware for handling Cross-Origin Resource Sharing
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// RequestLogger middleware for logging HTTP requests
func RequestLogger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})
}

// Recovery middleware for handling panics
func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered any) {
		if err, ok := recovered.(string); ok {
			log.Printf("Panic recovered: %s", err)
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal server error",
			"message": "Something went wrong",
		})
	})
}

// Rate limiting structure
type rateLimiter struct {
	requests map[string][]time.Time
	mutex    sync.RWMutex
	limit    int
	window   time.Duration
}

var (
	rateLimiters = make(map[string]*rateLimiter)
	rateMutex    sync.RWMutex
)

// RateLimit middleware for rate limiting requests
func RateLimit(limit int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		key := fmt.Sprintf("%s:%d:%v", clientIP, limit, window)

		rateMutex.Lock()
		limiter, exists := rateLimiters[key]
		if !exists {
			limiter = &rateLimiter{
				requests: make(map[string][]time.Time),
				limit:    limit,
				window:   window,
			}
			rateLimiters[key] = limiter
		}
		rateMutex.Unlock()

		limiter.mutex.Lock()
		now := time.Now()
		windowStart := now.Add(-window)

		// Clean old requests
		var validRequests []time.Time
		for _, reqTime := range limiter.requests[clientIP] {
			if reqTime.After(windowStart) {
				validRequests = append(validRequests, reqTime)
			}
		}
		limiter.requests[clientIP] = validRequests

		// Check if limit exceeded
		if len(validRequests) >= limit {
			limiter.mutex.Unlock()
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error":   "Rate limit exceeded",
				"message": fmt.Sprintf("Too many requests. Limit: %d per %v", limit, window),
			})
			c.Abort()
			return
		}

		// Add current request
		limiter.requests[clientIP] = append(limiter.requests[clientIP], now)
		limiter.mutex.Unlock()

		c.Next()
	}
}

// Cleanup function to remove old rate limiters
func cleanupRateLimiters() {
	ticker := time.NewTicker(1 * time.Hour)
	go func() {
		for range ticker.C {
			rateMutex.Lock()
			for key, limiter := range rateLimiters {
				limiter.mutex.Lock()
				now := time.Now()
				windowStart := now.Add(-limiter.window)

				// Check if all requests are old
				allOld := true
				for _, requests := range limiter.requests {
					for _, reqTime := range requests {
						if reqTime.After(windowStart) {
							allOld = false
							break
						}
					}
					if !allOld {
						break
					}
				}

				if allOld {
					delete(rateLimiters, key)
				}
				limiter.mutex.Unlock()
			}
			rateMutex.Unlock()
		}
	}()
}

// Initialize cleanup
func init() {
	cleanupRateLimiters()
}
