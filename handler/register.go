package handler

import (
	"easy-library/common"
	"easy-library/db"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	name := c.PostForm("user_name")           //将前端通过ajax传入的username的值传给name
	password := c.PostForm("password")        //将前端通过ajax传入的password的值传给password
	cusErr := db.UserRegister(name, password) //将name,password 传给UserRegister()方法
	if cusErr != nil {
		toHtml(c, cusErr, 0, nil)
	} else {
		toHtml(c, common.ErrSuccess, 1, nil)
	}
}
