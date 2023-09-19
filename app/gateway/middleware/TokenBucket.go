package middleware

import (
	"errors"
	"micro-todolist/pkg/ctl"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

type TokenBucket struct {
	capacity  int64      // 桶的容量
	rate      float64    // 令牌放入速率
	tokens    float64    // 当前令牌数量
	lastToken time.Time  // 上一次放令牌的时间
	mtx       sync.Mutex // 互斥锁
}

func (tb *TokenBucket) Allow() bool {
	tb.mtx.Lock()
	defer tb.mtx.Unlock()
	now := time.Now()
	// 计算需要放的令牌数量
	tb.tokens = tb.tokens + tb.rate*now.Sub(tb.lastToken).Seconds()
	if tb.tokens > float64(tb.capacity) {
		tb.tokens = float64(tb.capacity)
	}
	// 判断是否允许请求
	if tb.tokens >= 1 {
		tb.tokens--
		tb.lastToken = now
		return true
	} else {
		return false
	}
}

func LimitHandler(maxConn int) gin.HandlerFunc {
	bucket := ratelimit.NewBucketWithQuantum(time.Second, int64(maxConn), int64(maxConn))
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			c.JSON(http.StatusTooManyRequests, ctl.RespError(c, errors.New("too many request"), "Too many request", 429))
			c.Abort()
			return
		}
		c.Next()
	}
}
