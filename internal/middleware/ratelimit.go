package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
)

type RateLimitStrategy interface {
	Key(c *gin.Context) (string, error)
}

type IPRateLimit struct{}

func (s IPRateLimit) Key(c *gin.Context) (string, error) {
	ip := c.ClientIP()
	return "rate:" + ip + ":" + c.Request.Method + ":" + c.FullPath(), nil
}

type BusinessRateLimit struct{}

func (s BusinessRateLimit) Key(c *gin.Context) (string, error) {
	username := c.GetHeader("X-Username")
	apiKey := c.GetHeader("X-API-KEY")
	if apiKey == "" {
		return "", gin.Error{
			Err:  http.ErrNoCookie, // dummy sentinel
			Type: gin.ErrorTypeBind,
			Meta: "No API Key sent with the header",
		}
	}
	return "api_key:" + apiKey + ":rate:" + username + ":" + c.Request.Method + ":" + c.FullPath(), nil
}

func RateLimiter(rdb *redis.Client, limit int, window time.Duration, strat RateLimitStrategy) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		key, err := strat.Key(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		count, err := rdb.Incr(ctx, key).Result()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "rate limiter error," + err.Error()})
			return
		}
		if count == 1 {
			rdb.Expire(ctx, key, window)
		}
		if count > int64(limit) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
			return
		}
		c.Next()
	}
}
