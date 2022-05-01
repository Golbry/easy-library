package common

import "fmt"

type CusError struct {
	ErrNo   int32  `json:"errno"`
	Message string `json:"message"`
}

func NewCusError(errno int32, message string) *CusError {
	return &CusError{
		ErrNo:   errno,
		Message: message,
	}
}

func (c *CusError) Error() string {
	return fmt.Sprintf("{%d: %s}", c.ErrNo, c.Message)
}

func (c *CusError) GetErrNo() int32 {
	return c.ErrNo
}

func (c *CusError) GetMsg() string {
	return c.Message
}

var (
	ErrSuccess         = NewCusError(0, "success")
	ErrSystem          = NewCusError(1001, "系统错误")
	ErrInvalidUserName = NewCusError(1002, "无效的用户名")
	ErrExitedUserName  = NewCusError(1002, "已存在的用户名")
	ErrPassword        = NewCusError(1003, "密码错误")
)