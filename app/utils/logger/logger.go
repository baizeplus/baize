package logger

import (
	"baize/app/setting"
	"baize/app/utils/response"
	"fmt"
	"github.com/spf13/viper"

	"net/http"

	"os"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sl *zap.SugaredLogger

// Init 初始化lg
func init() {
	var core zapcore.Core
	var level zapcore.LevelEnabler
	var encoder zapcore.Encoder
	switch viper.GetString("log.level") {
	case "debug":
		level = zapcore.DebugLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}
	if viper.GetString("log.encoder") == "json" {
		encoder = getEncoder()
	} else {
		encoder = zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	}

	if viper.GetString("log.filename") == "" {
		core = zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), level)
	} else {
		writeSyncer := getLogWriter(viper.GetString("log.filename"), viper.GetInt("log.max_size"), viper.GetInt("log.max_backups"), viper.GetInt("log.max_age"))
		core = zapcore.NewCore(encoder, writeSyncer, level)
	}

	logger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger)
	sl = zap.L().Sugar()
	zap.L().Debug("init logger success")
	return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		cost := time.Since(start)
		if path != "/" {
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
	}
}

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func GinRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
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
		c.Next()
	}
}

type SqlyLog struct {
}

func (s *SqlyLog) Debug(cost time.Duration, sql string, args ...interface{}) {
	sl.Debug(sql, "\t", args, "\tcost:"+cost.String())
}
