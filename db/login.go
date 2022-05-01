package db

import (
	"database/sql"
	"easy-library/common"
	"easy-library/logger"
	"github.com/sirupsen/logrus"
)

// UserLogin /*登入账户*/
func UserLogin(username string, password string) (int, int8, int8, *common.CusError) {
	logHeader := "[UserLogin]"
	var destPassword string
	var status, level int8
	err := DB.QueryRow("SELECT password,status,level FROM user_info where user_name=?", username).Scan(&destPassword, &status, &level) //通过账户名查询密码
	if err != nil && err != sql.ErrNoRows {
		logger.DebugOutput.WithFields(logrus.Fields{
			"msg": logHeader,
		}).Warn("数据库查询失败", err, username, password)
		return 0, 0, 0, common.ErrSystem
	}
	if destPassword == password { //判断数据库密码 与前端输入的密码是否相同
		return 1, status, level, nil
	} else {
		return 0, 0, 0, common.ErrPassword
	}
}
