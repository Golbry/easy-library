package handler

import (
	"easy-library/db"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	name := c.PostForm("username")         //将前端通过ajax传入的username的值传给name
	password := c.PostForm("password")     //将前端通过ajax传入的password的值传给password
	result := db.UserLogin(name, password) //将name ,password传给UserLogin()方法
	toHtml(c, result)                      //返回 给html
	fmt.Println("ojbk")
}
