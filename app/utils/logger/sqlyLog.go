package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"regexp"
	"strings"
	"time"
)

type SqlyLog struct {
}

func (s *SqlyLog) Debug(cost time.Duration, sql string, args ...interface{}) {
	if zap.L().Level() == zapcore.DebugLevel {
		sql = strings.ReplaceAll(sql, "\n", " ")
		sql = strings.ReplaceAll(sql, "\t", " ")
		re := regexp.MustCompile(`\s+`)
		// 将多个空格替换成一个空格
		sql = re.ReplaceAllString(sql, " ")
		sql = strings.TrimSpace(sql)
		zap.L().Debug(sql, zap.Any("args", args), zap.Duration("cost", cost))
	}
}
