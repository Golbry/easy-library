package db

import (
	"easy-library/common"
	"easy-library/logger"
	"github.com/sirupsen/logrus"
)

func BorrowBook(bookId string, userName string, borrowNum uint) (int, *common.CusError) { //通过id查询  返回 books切片  通过server内的Querybook传入id
	logHeader := "[QueryBook]"
	tx, _ := DB.Begin()
	var leftSum, borrowedSum uint
	var bookName string
	err := tx.QueryRow("SELECT book_name,book_left_sum,book_borrowed_sum FROM book_global_info WHERE book_id = ?", bookId).Scan(&bookName, &leftSum, &borrowedSum) //执行查询语句 ,获取表中所有数据 并存入rows
	if err != nil {
		return 0, common.ErrSystem
	} //处理错误机制
	if borrowNum > leftSum {
		return 0, common.ErrInvalidBorrowNum
	}
	leftSum = leftSum - borrowNum
	borrowedSum = borrowedSum + borrowNum
	_, err = tx.Exec("UPDATE book_global_info SET book_left_sum=?,book_borrowed_sum=? WHERE book_id =?", leftSum, borrowedSum, bookId) //通过书籍id 修改 书籍时间 执行sql语句 将值传给results
	if err != nil {
		_ = tx.Rollback() //若失败需进行回滚
		logger.DebugOutput.WithFields(logrus.Fields{
			"msg": logHeader,
		}).Warn("数据库更新失败", err)
		return 0, common.ErrUpdate
	}
	_, err = tx.Exec("INSERT INTO book_history(book_id,book_name,status,borrower,borrow_num) VALUES(?,?,?,?,?)", bookId, bookName, 0, userName, borrowNum)
	if err != nil {
		_ = tx.Rollback()
		logger.DebugOutput.WithFields(logrus.Fields{
			"msg": logHeader,
		}).Warn("数据库插入失败", err)
		return 0, common.ErrUpdate
	}
	_ = tx.Commit()
	return 1, common.ErrSuccess
}
