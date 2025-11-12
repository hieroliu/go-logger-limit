package main

import (
	"math/rand"
	"time"

	"github.com/donnie4w/go-logger/logger"
	//"github.com/fatih/color"
	color "github.com/fatih/color"
)

// 通过 Formatter 构造格式输出
// 说明：
// {level} 日志级别数据
// {time}  日志时间
// {file}  日志文件
// {message}  日志主题内容

func main() {
	attrFormat := &logger.AttrFormat{
		SetLevelFmt: func(level logger.LEVELTYPE) string {
			switch level {
			case logger.LEVEL_DEBUG:
				return "DBG"
			case logger.LEVEL_INFO:
				return "INF"
			case logger.LEVEL_WARN:
				return "WRN"
			case logger.LEVEL_ERROR:
				return "ERR"
			case logger.LEVEL_FATAL:
				return "FTL"
			default:
				return "UNKNOWN"
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
				return []byte(color.HiGreenString(string(msg))) // Green for INFO
			case logger.LEVEL_WARN:
				return []byte(color.HiYellowString(string(msg))) // Yellow for WARN
			case logger.LEVEL_ERROR:
				return []byte(color.RedString(string(msg))) // Red for ERROR
			case logger.LEVEL_FATAL:
				//return []byte(color.New(color.BgRed, color.FgWhite).Sprint(string(msg))) // Red background for FATAL, then reset color
				return []byte(color.HiMagentaString(string(msg))) // Red background for FATAL, then reset color
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

	formatter := "[{level} {time}] {message}\n"

	logger.SetOption(&logger.Option{
		Console:    true,
		Formatter:  formatter,
		FileOption: fileOption,
		AttrFormat: attrFormat})
	//logger.SetFormatter("[{level} {time}] {message}\n")
	logger.Debug("1 this is a debug message")
	logger.Info("2 this is a info message")
	logger.Warn("3 this is a warn message")
	logger.Error("4 this is a error message")
	logger.Fatal("5 this is a fatal message") // 移除Fatal调用，避免程序退出
	for i := 0; i < 1000; i++ {
		randLevel := logger.LEVELTYPE(rand.Intn(6) + 1) // 1-5，避开FATAL级别
		switch randLevel {
		case logger.LEVEL_DEBUG:
			logger.Debugf("%d %s %d", i, "this is a debug message", int16(randLevel))
		case logger.LEVEL_INFO:
			logger.Infof("%d %s %d", i, "this is a info message", int16(randLevel))
		case logger.LEVEL_WARN:
			logger.Warnf("%d %s %d", i, "this is a warn message", int16(randLevel))
		case logger.LEVEL_ERROR:
			logger.Errorf("%d %s %d", i, "this is a error message", int16(randLevel))
		case logger.LEVEL_FATAL:
			logger.Fatalf("%d %s %d", i, "this is a fatal message", int16(randLevel))
		default:
			logger.Infof("default this is a info message %d", int16(randLevel))
		}
	}
	logger.Debug("this is a debug message")
	logger.Info("this is a info message")
	logger.Warn("this is a warn message")
	logger.Error("this is a error message")
	logger.Fatal("this is a fatal message") // 移除Fatal调用，避免程序退出

}
