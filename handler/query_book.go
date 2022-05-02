package handler

import (
	"easy-library/common"
	"easy-library/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func QueryBook(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id")) //将前端通过ajax传入的id的值传给id
	fmt.Println(c.PostForm("id"))
	if err != nil {
		panic(err)
	} //处理错误机制
	results := db.QueryBook(id) //将id  传给QueryBooks()方法
	fmt.Println(results)
	toHtml(c, common.ErrSuccess, results, nil) //将results返回给html
}
