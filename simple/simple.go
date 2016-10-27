package main

import "github.com/lucky2me/log"

func main() {
	// 初始化一个logger, 并传入一个日志的绝对路径
	logger := log.NewLogger("/home/golang/src/github.com/lucky2me/log/simple/logs", log.LoggerLevelInfo)

	logger.Error("error test")
	logger.Info("info test")
	logger.Debug("debug test")
}
