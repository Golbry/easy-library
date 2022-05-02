package handler

import (
	"easy-library/db"
	"github.com/gin-gonic/gin"
)

func QueryBook(c *gin.Context) {
	bookName := c.PostForm("book_name")       //将前端通过ajax传入的id的值传给id
	results, cusErr := db.QueryBook(bookName) //将id  传给QueryBooks()方法
	toHtml(c, cusErr, results, nil)           //将results返回给html
}
