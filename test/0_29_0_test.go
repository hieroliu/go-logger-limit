package test

import (
	"testing"

	"github.com/hieroliu/go-logger-limit/logger"
)

// 通过 Formatter 构造json格式输出
// 说明：
// {level} 日志级别数据
// {time}  日志时间
// {file}  日志文件
// {message}  日志主题内容
func Test_formater(t *testing.T) {
	logger.SetOption(&logger.Option{Console: true, Formatter: `{"level":"{level}","time":"{time}","file":"{file}","message":"{message}"}` + "\n"})
	logger.Debug("this is a debug message")
	logger.Info("this is a info message")
	logger.Warn("this is a warn message")
	logger.Error("this is a error message")
}

// 自定义json格式的参数
// 比如修改LEVEL的格式，同理，可以修改时间的格式
func Test_formater2(t *testing.T) {
	levelFmt := func(level logger.LEVELTYPE) string {
		switch level {
		case logger.LEVEL_DEBUG:
			return "debug"
		case logger.LEVEL_INFO:
			return "info"
		case logger.LEVEL_FATAL:
			return "fatal"
		case logger.LEVEL_WARN:
			return "warn"
		case logger.LEVEL_ERROR:
			return "error"
		default:
			return "unknown"
		}
	}
	logger.SetOption(&logger.Option{Console: true, Format: logger.FORMAT_LEVELFLAG | logger.FORMAT_DATE | logger.FORMAT_TIME, Formatter: `{"level":"{level}","time":"{time}","message":"{message}"}` + "\n", AttrFormat: &logger.AttrFormat{SetLevelFmt: levelFmt}})
	logger.Debug("this is a debug message")
	logger.Info("this is a info message")
	logger.Warn("this is a warn message")
	logger.Error("this is a error message")
}
