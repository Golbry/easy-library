package db

import (
	"easy-library/common"
	"easy-library/logger"
	"github.com/sirupsen/logrus"
)

func ReturnBook(bookId string, userName string, returnNum uint) (int, *common.CusError) { //通过id查询  返回 books切片  通过server内的Querybook传入id
	logHeader := "[QueryBook]"
	tx, _ := DB.Begin()
	var borrowNum uint
	err := tx.QueryRow("SELECT borrow_num FROM book_history WHERE book_id = ? AND borrower=?", bookId, userName).Scan(&borrowNum)
	if err != nil {
		return 0, common.ErrSystem
	} //处理错误机制
	if borrowNum < returnNum {
		return 0, common.ErrInvalidBorrowNum
	}
	borrowNum = borrowNum - returnNum
	_, err = tx.Exec("UPDATE book_history SET borrow_num=? WHERE book_id =? AND borrower=?", borrowNum, bookId, userName)
	if err != nil {
		_ = tx.Rollback() //若失败需进行回滚
		logger.DebugOutput.WithFields(logrus.Fields{
			"msg": logHeader,
		}).Warn("数据库更新失败", err)
		return 0, common.ErrUpdate
	}
	var leftSum, borrowedSum uint
	err = tx.QueryRow("SELECT book_left_sum,book_borrowed_sum FROM book_global_info WHERE book_id = ?", bookId).Scan(&leftSum, &borrowedSum)
	if err != nil {
		return 0, common.ErrSystem
	} //处理错误机制
	leftSum = leftSum + returnNum
	borrowedSum = borrowedSum - returnNum
	_, err = tx.Exec("UPDATE book_global_info SET book_left_sum=?,book_borrowed_sum=? WHERE book_id =?", leftSum, borrowedSum, bookId) //通过书籍id 修改 书籍时间 执行sql语句 将值传给results
	if err != nil {
		_ = tx.Rollback() //若失败需进行回滚
		logger.DebugOutput.WithFields(logrus.Fields{
			"msg": logHeader,
		}).Warn("数据库更新失败", err)
		return 0, common.ErrUpdate
	}
	_ = tx.Commit()
	return 1, common.ErrSuccess
}
