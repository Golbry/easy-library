package db

import (
	"easy-library/common"
	"easy-library/logger"
	"github.com/sirupsen/logrus"
)

func QueryBook(bookName string) ([]BookInfo, *common.CusError) { //通过id查询  返回 books切片  通过server内的Querybook传入id
	logHeader := "[QueryBook]"
	books := make([]BookInfo, 0)                                                                                                                              //创建空的users切片
	rows, err := DB.Query("SELECT book_id,book_name,author,press,book_left_sum,book_borrowed_sum FROM book_global_info WHERE book_name LIKE ?", bookName+"%") //执行查询语句 ,获取表中所有数据 并存入rows
	if err != nil {
		return nil, common.ErrSystem
	} //处理错误机制
	var book BookInfo //将结构体books 定义为book
	for rows.Next() { //for循环 rows里的数据
		err = rows.Scan(&book.BookId, &book.BookName, &book.Author, &book.Press, &book.BookLeftSum, &book.BookBorrowedSum) //扫描数据
		books = append(books, book)                                                                                        //将book结构体内的数据传入到users切片当中
		if err != nil {
			logger.DebugOutput.WithFields(logrus.Fields{
				"msg": logHeader,
			}).Warn("数据库查询失败", err)
			return nil, common.ErrQuery
		}
		logger.DebugOutput.WithFields(logrus.Fields{
			"msg": logHeader,
		}).Warn(book)
	}
	return books, common.ErrSuccess
}
