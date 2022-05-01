package handler

import (
	"easy-library/common"
	"easy-library/db"
	"easy-library/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Login(c *gin.Context) {
	logHeader := "[Login]"
	name := c.PostForm("user_name")                               //将前端通过ajax传入的username的值传给name
	password := c.PostForm("password")                            //将前端通过ajax传入的password的值传给password
	result, status, level, cusErr := db.UserLogin(name, password) //将name ,password传给UserLogin()方法
	if cusErr != nil {
		logger.DebugOutput.WithFields(logrus.Fields{
			"msg": logHeader,
		}).Error(cusErr, name, password)
		toHtml(c, cusErr, result, nil)
	} else {
		toHtml(c, common.ErrSuccess, result, map[string]int8{
			"status": status,
			"level":  level,
		})
	}
}
