package main

import (
	"easy-library/db"
	"easy-library/handler"
	"easy-library/id_generator"
	"easy-library/logger"
	"fmt"
)

func main() {
	logger.Init()
	id_generator.Init()
	fmt.Println(id_generator.GenerateId())
	db.MustInit() //初始化数据库
	handler.Run() //注册并运行接口
}
