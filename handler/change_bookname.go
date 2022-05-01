package handler

import (
	"easy-library/common"
	"easy-library/db"
	"github.com/gin-gonic/gin"
	"strconv"
)

func changeBookname(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id")) //将前端传入的数据赋值给id  (需类型转换)
	if err != nil {
		panic(err)
	} //处理错误机制
	bookname := c.PostForm("bookName")    //将前端传入的数据赋值给bookname
	re := db.ChangBookName(id, bookname)  //将id ,bookname 传给ChangBookName()方法
	toHtml(c, common.ErrSuccess, re, nil) //返回给html
}
