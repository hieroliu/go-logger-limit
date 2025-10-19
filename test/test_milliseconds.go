package main

import (
	"time"
	"github.com/hieroliu/go-logger-limit/logger"
)

func main() {
	// 重置默认设置
	logger.SetFormatter("{level}{time} {message}\n")
	
	// 1. 测试毫秒格式选项
	logger.SetFormat(logger.FORMAT_LEVELFLAG | logger.FORMAT_DATE | logger.FORMAT_TIME | logger.FORMAT_MILLISECONDS)
	logger.Info("毫秒格式测试：应该显示3位小数")
	
	// 短暂延迟
	time.Sleep(10 * time.Millisecond)
	
	// 2. 对比微秒格式
	logger.SetFormat(logger.FORMAT_LEVELFLAG | logger.FORMAT_DATE | logger.FORMAT_TIME | logger.FORMAT_MICROSECONDS)
	logger.Info("微秒格式测试：应该显示6位小数")
	
	// 短暂延迟
	time.Sleep(10 * time.Millisecond)
	
	// 3. 自定义格式测试，包含级别
	logger.SetFormatter("{time} - {level} - {message}\n")
	logger.SetFormat(logger.FORMAT_LEVELFLAG | logger.FORMAT_DATE | logger.FORMAT_TIME | logger.FORMAT_MILLISECONDS)
	logger.Info("自定义格式+毫秒测试")
}