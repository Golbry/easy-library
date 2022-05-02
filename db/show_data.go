package db

import (
	"easy-library/common"
	"easy-library/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

//定义书籍结构体
type BookInfo struct { // 结果集，参数名需大写
	BookId          string `json:"book_id"`
	BookName        string `json:"book_name"`
	Author          string `json:"author"`
	Press           string `json:"press"`
	BookLeftSum     uint   `json:"book_left_sum"`
	BookBorrowedSum uint   `json:"book_borrowed_sum"`
}

// Showdata 动态数据
func ShowData() ([]BookInfo, *common.CusError) { //返回切片数据
	logHeader := "[ShowData]"
	books := make([]BookInfo, 0)                                                                                         //创建空的users切片
	rows, err := DB.Query("SELECT book_id,book_name,author,press,book_left_sum,book_borrowed_sum FROM book_global_info") //执行查询语句 ,获取表中所有数据 并存入rows
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
