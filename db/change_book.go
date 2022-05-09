package db

import (
	"easy-library/common"
	"easy-library/id_generator"
	"easy-library/logger"
	"fmt"
	"github.com/sirupsen/logrus"
)

//修改书籍时间信息
func ChangeBook(bookId string, bookName string, author string, press string, bookLeftSum uint, status int8) (int, *common.CusError) { //前端传入的书籍 id 与 书籍时间 作为参数     //通过server内的changeBooktime传入id bookTime
	preStr := fmt.Sprintf("%s%s%s", bookName, author, press) //名称、作者和出版社唯一确定一本书，用它们来计算sha256，可确保图书的唯一性
	bookIdChanged := id_generator.GenerateId(preStr)
	_, err := DB.Exec("UPDATE book_global_info SET book_id=?,book_name=?,author=?,press=?,book_left_sum=?,status=? where book_id =?", bookIdChanged, bookName, author, press, bookLeftSum, status, bookId) //通过书籍id 修改 书籍时间 执行sql语句 将值传给results
	if err != nil {
		logger.DebugOutput.WithFields(logrus.Fields{
			"msg": "logHeader",
		}).Error(bookId, status, err)
		return 0, common.ErrUpdate
	} else {
		return 1, common.ErrSuccess
	}
}
