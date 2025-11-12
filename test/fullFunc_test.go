package test

import (
	"testing"
	"time"

	"github.com/donnie4w/go-logger/logger"
	color "github.com/fatih/color"
)

// 通过 Formatter 构造格式输出
// 说明：
// {level} 日志级别数据
// {time}  日志时间
// {file}  日志文件
// {message}  日志主题内容

func Test_fullFunc_Formatter(t *testing.T) {
	attrFormat := &logger.AttrFormat{
		SetLevelFmt: func(level logger.LEVELTYPE) string {
			switch level {
			case logger.LEVEL_DEBUG:
				return "[DBG]"
			case logger.LEVEL_INFO:
				return "[INF]"
			case logger.LEVEL_WARN:
				return "[WRN]"
			case logger.LEVEL_ERROR:
				return "[ERR]"
			case logger.LEVEL_FATAL:
				return "[FTL]"
			default:
				return "[UNKNOWN]"
			}
		},
		SetTimeFmt: func() (string, string, string) {
			now := time.Now().Format("15:04:05.000")
			return now, "", ""
		},
		SetConsoleColorFmt: func(level logger.LEVELTYPE, msg []byte) []byte {
			switch level {
			case logger.LEVEL_DEBUG:
				return []byte(color.HiCyanString(string(msg))) // Blue for DEBUG
			case logger.LEVEL_INFO:
				return []byte(color.GreenString(string(msg))) // Green for INFO
			case logger.LEVEL_WARN:
				return []byte(color.HiYellowString(string(msg))) // Yellow for WARN
			case logger.LEVEL_ERROR:
				return []byte(color.RedString(string(msg))) // Red for ERROR
			case logger.LEVEL_FATAL:
				return []byte(color.New(color.BgRed, color.FgWhite).Sprint(string(msg))) // Red background for FATAL
			default:
				return msg
			}

		},
	}

	fileOption := &logger.FileMixedMode{
		Filename:   "testFullmixid.log",
		Maxbackup:  10,
		IsCompress: true,
		Timemode:   logger.MODE_DAY,
		Maxsize:    1 << 20,
	}

	formatter := "{time} {level} {message}\n"

	logger.SetOption(&logger.Option{
		Console:    true,
		Formatter:  formatter,
		FileOption: fileOption,
		AttrFormat: attrFormat})
	//logger.SetFormatter("{level}{time} {message}\n")
	logger.Debug("this is a debug message")
	logger.Info("this is a info message")
	logger.Warn("this is a warn message")
	logger.Error("this is a error message")
}
