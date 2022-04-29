package handler

import (
	"easy-library/db"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Addbook(c *gin.Context) {
	name := c.PostForm("bookname")        //将前端传入的数据赋值给name
	time := c.PostForm("goodShelfStatus") //将前端传入的数据赋值给time
	re := db.Add(name, time)              //将name,time 传给Add()方法
	fmt.Println(re)
	toHtml(c, re) //返回给html文件
}
