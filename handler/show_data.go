package handler

import (
	"easy-library/db"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Showdata(c *gin.Context) {
	a := db.Showdata() //将showdata()获取的值 赋值给a
	//fmt.Println(a, m, d)
	fmt.Println(a)
	toHtml(c, a) //将a 返回给html
}
