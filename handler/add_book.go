package handler

import (
	"easy-library/common"
	"easy-library/db"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AddBook(c *gin.Context) {
	bookName := c.PostForm("book_name")
	author := c.PostForm("author")
	press := c.PostForm("press")
	bookSum := c.PostForm("book_left_sum")
	if len(bookName) > common.MaxStrLen || len(author) > common.MaxStrLen ||
		len(author) > common.MaxStrLen || len(bookSum) > common.MaxStrLen {
		toHtml(c, common.ErrInvalidStr, 0, nil)
	}
	intNum, _ := strconv.Atoi(bookSum)
	if intNum <= 0 {
		toHtml(c, common.ErrInvalidNum, 0, nil)
	}
	bookSumUint := uint(intNum)                                 //格式转换
	res, cusErr := db.Add(bookName, author, press, bookSumUint) //将name,time 传给Add()方法
	toHtml(c, cusErr, res, nil)                                 //返回给html文件
}
