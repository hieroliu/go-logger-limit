package main

import (
	"time"

	"github.com/hieroliu/go-logger-limit/logger"
)

// 注意：在实际使用中，应该使用模块路径导入，而不是绝对路径
// 这里为了测试目的暂时使用绝对路径

func main() {
	// 设置日志格式，包含级别标签、日期、时间和毫秒
	logger.SetFormat(logger.FORMAT_LEVELFLAG | logger.FORMAT_DATE | logger.FORMAT_TIME | logger.FORMAT_MILLISECONDS)

	// 设置日志级别为DEBUG，显示所有级别的日志
	logger.SetLevel(logger.LEVEL_DEBUG)

	// 测试各种级别的日志输出，验证不同颜色
	logger.Debug("这是一条DEBUG级别的日志，应该显示为蓝色")
	logger.Info("这是一条INFO级别的日志，应该显示为绿色")
	logger.Warn("这是一条WARN级别的日志，应该显示为黄色")
	logger.Error("这是一条ERROR级别的日志，应该显示为红色")
	logger.Fatal("这是一条FATAL级别的日志，应该显示为红色背景")

	// 注意：FATAL级别的日志会导致程序退出
	// 为了不影响测试，我们先暂停一下，让前面的日志输出完成
	time.Sleep(100 * time.Millisecond)

	// 可选：取消下面这行的注释来测试FATAL级别（会导致程序退出）
	// logger.Fatal("这是一条FATAL级别的日志，应该显示为红色背景")
}
