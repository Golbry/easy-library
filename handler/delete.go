package handler

import (
	"easy-library/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Delect(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id")) //将前端传入的数据赋值给id  (需类型转换)
	fmt.Println(c.PostForm("id"))
	if err != nil {
		panic(err) //处理错误机制
	}
	result := db.Delect(id) //将id 传给Delect()方法
	toHtml(c, result)       //返回html
}
