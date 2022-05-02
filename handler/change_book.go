package handler

import (
	"easy-library/common"
	"easy-library/db"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ChangeBook(c *gin.Context) {
	bookId := c.PostForm("book_id")
	bookName := c.PostForm("book_name")
	author := c.PostForm("author")
	press := c.PostForm("press")
	status := c.PostForm("status")
	bookLeftSum := c.PostForm("book_left_sum")
	if len(bookName) > common.MaxStrLen || len(author) > common.MaxStrLen ||
		len(author) > common.MaxStrLen || len(bookLeftSum) > common.MaxStrLen || len(bookId) > common.MaxStrLen ||
		len(status) > common.MaxStrLen {
		toHtml(c, common.ErrInvalidStr, 0, nil)
	}
	intNum, _ := strconv.Atoi(bookLeftSum)
	if intNum <= 0 {
		toHtml(c, common.ErrInvalidNum, 0, nil)
	}
	bookSumUint := uint(intNum) //格式转换
	intNum, _ = strconv.Atoi(status)
	if intNum < 0 || intNum > 127 {
		toHtml(c, common.ErrInvalidNum, 0, nil)
	}
	bookStatus := int8(intNum)
	res, cusErr := db.ChangeBook(bookId, bookName, author, press, bookSumUint, bookStatus) //将name,time 传给Add()方法
	toHtml(c, cusErr, res, nil)
}
