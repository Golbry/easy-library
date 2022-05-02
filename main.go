package main

import (
	"easy-library/db"
	"easy-library/handler"
	"easy-library/logger"
)

func main() {
	logger.Init()
	db.MustInit() //初始化数据库
	handler.Run() //注册并运行接口
}
