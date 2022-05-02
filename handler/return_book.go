package handler

import (
	"easy-library/common"
	"easy-library/db"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ReturnBook(c *gin.Context) {
	bookId := c.PostForm("book_id")          //将前端通过ajax传入的id的值传给id
	userName := c.PostForm("user_name")      //用户名
	returnNumStr := c.PostForm("return_num") //归还数量
	if len(bookId) > common.MaxStrLen || len(userName) > common.MaxStrLen ||
		len(returnNumStr) > common.MaxStrLen {
		toHtml(c, common.ErrInvalidStr, 0, nil)
	}
	intNum, _ := strconv.Atoi(returnNumStr)
	if intNum <= 0 {
		toHtml(c, common.ErrInvalidNum, 0, nil)
	}
	borrowNum := uint(intNum) //格式转换
	results, cusErr := db.ReturnBook(bookId, userName, borrowNum)
	toHtml(c, cusErr, results, nil)
}
