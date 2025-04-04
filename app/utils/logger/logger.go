package logger

import (
	"baize/app/setting"
	"fmt"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"runtime/debug"
)

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

func GoroutineRecovery(path string, any any) {
	zap.L().Error("goroutine",
		zap.Any("error", any),
		zap.String("path", path),
		zap.String("stack", string(debug.Stack())),
	)
	if setting.Conf.Mode == "dev" {
		fmt.Println("----------------------------------------------------------------------------------------------------")
		fmt.Printf("error:%s\n", any)
		fmt.Println("stack:" + string(debug.Stack()))
	}
}
