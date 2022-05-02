package db

import (
	"easy-library/common"
	"easy-library/id_generator"
	"fmt"
)

// Delect 删除书籍
func Delect(id int) int { //通过 前端传入的id 数值 参数     //通过server内的Delect传入id
	var a int
	results, err := DB.Exec("DELETE FROM book_wyb where id=?", id) // 通过id 执行 删除sql语句 删除数据
	if err != nil {
		panic(err)
	} else { //处理错误机制
		fmt.Println(results)
		a = 1
	}
	return a
}

// ChangBookName 修改书籍名称信息
func ChangBookName(id int, bookname string) int { //前端传入的书籍 id 与 书籍名称 作为参数   //通过server内的changeBookname传入id bookname
	var a int
	results, err := DB.Exec("update book_wyb set bookName =? where id= ? ", bookname, id) //通过书籍id 修改 书籍名称 执行sql语句 将值传给results
	if err != nil {
		panic(err)
	} else { //执行错误机制
		fmt.Println(results)
		a = 1
	}
	return a
}

//修改书籍时间信息
func ChangeBook(bookId string, bookName string, author string, press string, bookLeftSum uint, status int8) (int, *common.CusError) { //前端传入的书籍 id 与 书籍时间 作为参数     //通过server内的changeBooktime传入id bookTime
	preStr := fmt.Sprintf("%s%s%s", bookName, author, press) //名称、作者和出版社唯一确定一本书，用它们来计算sha256，可确保图书的唯一性
	bookIdChanged := id_generator.GenerateId(preStr)
	_, err := DB.Exec("UPDATE book_global_info SET book_id=?,book_name=?,author=?,press=?,book_left_sum=?,status=? where book_id =?", bookIdChanged, bookName, author, press, bookLeftSum, status, bookId) //通过书籍id 修改 书籍时间 执行sql语句 将值传给results
	if err != nil {
		return 0, common.ErrUpdate
	} else {
		return 1, common.ErrSuccess
	}
}
