package handler

import (
	"easy-library/common"
	"easy-library/db"
	"github.com/gin-gonic/gin"
	"strconv"
)

func BorrowBook(c *gin.Context) {
	bookId := c.PostForm("book_id")          //将前端通过ajax传入的id的值传给id
	userName := c.PostForm("user_name")      //用户名
	borrowNumStr := c.PostForm("borrow_num") //借阅数量
	if len(bookId) > common.MaxStrLen || len(userName) > common.MaxStrLen ||
		len(borrowNumStr) > common.MaxStrLen {
		toHtml(c, common.ErrInvalidStr, 0, nil)
	}
	intNum, _ := strconv.Atoi(borrowNumStr)
	if intNum <= 0 {
		toHtml(c, common.ErrInvalidNum, 0, nil)
	}
	borrowNum := uint(intNum)                                     //格式转换
	results, cusErr := db.BorrowBook(bookId, userName, borrowNum) //将id  传给QueryBooks()方法
	toHtml(c, cusErr, results, nil)                               //将results返回给html
}
