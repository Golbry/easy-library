package handler

import (
	"easy-library/db"
	"easy-library/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ShowData(c *gin.Context) {
	res, cusErr := db.ShowData() //将showdata()获取的值 赋值给a
	logger.DebugOutput.WithFields(logrus.Fields{
		"msg": "logHeader",
	}).Info(res)
	toHtml(c, cusErr, res, nil) //将a 返回给html
}
