package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var DebugOutput *logrus.Logger

func Init() {
	DebugOutput = logrus.New()
	DebugOutput.Formatter = &logrus.JSONFormatter{} // 设置为json格式的日志
	gin.SetMode(gin.DebugMode)                      // 调试模式
	DebugOutput.Level = logrus.InfoLevel            // 设置日志级别
}
