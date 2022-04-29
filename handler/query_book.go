package handler

import (
	"easy-library/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Querybook(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id")) //将前端通过ajax传入的id的值传给id
	fmt.Println(c.PostForm("id"))
	if err != nil {
		panic(err)
	} //处理错误机制
	results := db.Querybooks(id) //将id  传给Querybooks()方法
	fmt.Println(results)
	toHtml(c, results) //将results返回给html
}
