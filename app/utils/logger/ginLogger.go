package logger

import (
	"baize/app/baize"
	"baize/app/setting"
	"baize/app/utils/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"runtime/debug"
	"time"
)

type loggerMiddlewareBuilder struct {
	paths baize.Set[string]
}

func NewLoggerMiddlewareBuilder() *loggerMiddlewareBuilder {
	return &loggerMiddlewareBuilder{
		paths: baize.Set[string]{},
	}
}

func (l *loggerMiddlewareBuilder) IgnorePaths(path string) *loggerMiddlewareBuilder {
	l.paths.Add(path)
	return l
}

func (l *loggerMiddlewareBuilder) Build() func(c *gin.Context) {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		defer func() {
			if err := recover(); err != nil {
				zap.L().Error("[Recovery from panic]",
					zap.Any("error", err),
					zap.String("path", c.Request.URL.Path),
					zap.String("query", c.Request.URL.RawQuery),
					zap.String("ip", c.ClientIP()),
					zap.String("user-agent", c.Request.UserAgent()),
					zap.String("stack", string(debug.Stack())),
				)
				if setting.Conf.Mode == "dev" {
					fmt.Println("----------------------------------------------------------------------------------------------------")
					fmt.Printf("error:%s\n", err)
					fmt.Println("stack:" + string(debug.Stack()))
				}
				c.JSON(http.StatusInternalServerError, response.ResponseData{Code: response.Error, Msg: response.Error.Msg()})
			}
		}()

		defer func() {
			cost := time.Since(start)
			if c.Request.Method != "OPTIONS" && !l.paths.Contains(c.Request.RequestURI) {
				zap.L().Info(path,
					zap.Int("status", c.Writer.Status()),
					zap.String("method", c.Request.Method),
					zap.String("path", path),
					zap.String("query", query),
					zap.String("ip", c.ClientIP()),
					zap.String("user-agent", c.Request.UserAgent()),
					zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
					zap.Duration("cost", cost),
				)
			}
		}()
		c.Next()
	}
}
