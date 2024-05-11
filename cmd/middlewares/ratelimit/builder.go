package ratelimit

import (
	_ "embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Builder struct {
	prefix   string
	cmd      redis.Cmdable
	interval time.Duration
	rate     int
}

//go:embed slide_window.lua
var luaScript string

func NewBuilder(cmd redis.Cmdable, interval time.Duration, rate int) *Builder {
	return &Builder{
		cmd:      cmd,
		prefix:   "ip-limit",
		interval: interval,
		rate:     rate,
	}
}

func (b *Builder) Prefix(prefix string) *Builder {
	b.prefix = prefix
	return b
}

func (b *Builder) Build() gin.HandlerFunc {
	return func(c *gin.Context) {
		limited, err := b.limit(c)
		if err != nil {
			zap.L().Error("Build Redis Fail ", zap.Error(err))
			// Redis 崩溃了 防止全打到数据库 限流
			c.AbortWithStatus(http.StatusInternalServerError)
			// 不限流 可能会打崩数据库
			// ctx.Next()
			return
		}
		if limited {
			zap.L().Error("Build 触发限流 ", zap.Error(err))
			c.AbortWithStatus(http.StatusTooManyRequests)
			return
		}
		c.Next()
	}
}

func (b *Builder) limit(c *gin.Context) (bool, error) {
	key := fmt.Sprintf("%s:%s", b.prefix, c.ClientIP())
	return b.cmd.Eval(c, luaScript, []string{key}, b.interval.Milliseconds(), b.rate, time.Now().UnixMilli()).Bool()
}
