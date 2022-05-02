package db

import (
	"easy-library/common"
	"easy-library/id_generator"
	"fmt"
)

// Add 添加书籍
func Add(bookname string, author string, press string, currentNum uint) (int, *common.CusError) { //通过书籍名称
	preStr := fmt.Sprintf("%s%s%s", bookname, author, press) //名称、作者和出版社唯一确定一本书，用它们来计算sha256，可确保图书的唯一性
	bookId := id_generator.GenerateId(preStr)
	_, err := DB.Exec("INSERT INTO book_global_info(book_id,book_name,author,press,book_left_sum) VALUES(?,?,?,?,?)", bookId, bookname, author, press, currentNum) //执行 添加sql语句
	if err != nil {
		return 0, common.ErrExistedBook
	} else { //处理错误机制
		return 1, common.ErrSuccess
	}
}
