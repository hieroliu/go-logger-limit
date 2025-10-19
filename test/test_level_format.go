package main

import (
	"time"
	"github.com/hieroliu/go-logger-limit/logger"
)

func main() {
	// 设置日志格式
	logger.SetFormat(logger.FORMAT_LEVELFLAG | logger.FORMAT_DATE | logger.FORMAT_TIME | logger.FORMAT_MILLISECONDS)
	
	// 测试所有日志级别
	logger.Debug("这是一条DEBUG级别日志")
	logger.Info("这是一条INFO级别日志")
	logger.Warn("这是一条WARN级别日志")
	logger.Error("这是一条ERROR级别日志")
	logger.Fatal("这是一条FATAL级别日志")
	
	// 等待日志输出完成
	time.Sleep(100 * time.Millisecond)
}