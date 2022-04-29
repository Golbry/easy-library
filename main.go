package main

import (
	"easy-library/db"
	"easy-library/handler"
)

func main() {
	db.MustInit() //初始化数据库
	handler.Run() //注册并运行接口
}
