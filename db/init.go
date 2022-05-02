package db

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

//数据库连接
func MustInit() {
	var err error
	//"用户名:密码@[连接方式](主机名:端口号)/数据库名"
	DB, err = sql.Open("mysql", "easy_library:McFc2jLb6Am8PkYz@tcp(81.70.240.235:3306)/easy_library?charset=utf8") // 设置连接数据库的参数
	if err != nil {
		panic(err)
	}
	err = DB.Ping() //连接数据库
	if err != nil {
		fmt.Println("数据库连接失败")
	}
}
