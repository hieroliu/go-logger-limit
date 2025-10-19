package main

import (
	"github.com/hieroliu/go-logger-limit/logger"
)

func main() {
	// 设置简单的格式，只包含级别、时间和消息
	logger.SetFormatter("{level} {time} {message}\n")

	// 单独测试毫秒格式 - 这是我们新增的功能
	logger.SetFormat(logger.FORMAT_LEVELFLAG | logger.FORMAT_DATE | logger.FORMAT_TIME | logger.FORMAT_MILLISECONDS)
	logger.Info("毫秒格式测试 (应该显示3位小数)")

	// 对比测试微秒格式
	logger.SetFormat(logger.FORMAT_LEVELFLAG | logger.FORMAT_DATE | logger.FORMAT_TIME | logger.FORMAT_MICROSECONDS)
	logger.Info("微秒格式测试 (应该显示6位小数)")

	// 对比测试微秒格式
	logger.SetFormat(logger.FORMAT_LEVELFLAG | logger.FORMAT_TIME | logger.FORMAT_MICROSECONDS)
	logger.Info("微秒格式测试 (应该显示6位小数)")

	// 单独测试毫秒格式
	logger.SetFormat(logger.FORMAT_LEVELFLAG | logger.FORMAT_TIME | logger.FORMAT_MILLISECONDS)
	logger.Info("微秒格式测试 (应该显示3位小数)")
}
