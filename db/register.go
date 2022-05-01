package db

import (
	"database/sql"
	"easy-library/common"
	"easy-library/logger"
	"github.com/sirupsen/logrus"
)

// UserRegister 用户注册
func UserRegister(username string, password string) *common.CusError {
	logHeader := "[UserRegister]"
	var destUserName string
	err := DB.QueryRow("SELECT user_name FROM user_info WHERE user_name=?", username).Scan(&destUserName) //查询user_info内所有信息
	if err != nil && err != sql.ErrNoRows {
		logger.DebugOutput.WithFields(logrus.Fields{
			"msg": logHeader,
		}).Warn("数据库查询失败", err)
		return common.ErrSystem
	}
	if username != destUserName || err == sql.ErrNoRows { //判断数据库内的u.name是否与前端页面传入的username相匹配 若不匹配 则无此用户
		results, err := DB.Exec("INSERT INTO user_info(user_name,password) VALUES(?,?)", username, password) //执行插入sql语句
		if err != nil {
			logger.DebugOutput.WithFields(logrus.Fields{
				"msg": logHeader,
			}).Error("执行失败") //处理错误机制
			return common.ErrSystem
		} else {
			rows, _ := results.RowsAffected() //输出执行的行数
			if rows != 1 {
				logger.DebugOutput.WithFields(logrus.Fields{
					"msg": logHeader,
				}).Error("执行失败") //处理错误机制
				return common.ErrSystem
			} else {
				return nil //执行成功
			}
		}
	} else {
		return common.ErrExitedUserName
	}
}
