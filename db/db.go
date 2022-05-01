package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//定义书籍结构体
type books_wyb struct { // 结果集，参数名需大写
	Bookid    int
	BookName  string
	CreatTime string
}

// Showdata 动态数据
func Showdata() []books_wyb { //返回切片数据
	users := make([]books_wyb, 0)                     //创建空的users切片
	rows, err := DB.Query("SELECT  *  from book_wyb") //执行查询语句 ,获取表中所有数据 并存入rows
	if err != nil {
		panic(err)
	} //处理错误机制
	var book books_wyb //将结构体books 定义为book
	for rows.Next() {  //for循环 rows里的数据
		err = rows.Scan(&book.Bookid, &book.BookName, &book.CreatTime) //扫描数据
		users = append(users, book)                                    //将book结构体内的数据传入到users切片当中
		if err != nil {
			fmt.Println(err) //错误处理机制
		}
	}
	fmt.Println("aa", users)
	return users
}

// Querybooks 查询书籍
func Querybooks(id int) []books_wyb { //通过id查询  返回 books切片  通过server内的Querybook传入id
	users := make([]books_wyb, 0)                                    //创建 users空切片
	rows, err := DB.Query("SELECT  *  from book_wyb where id=?", id) //通过id执行查询 sql语句
	if err != nil {
		panic(err)
	} //执行错误机制
	var book books_wyb //将结构体 定义为 book
	for rows.Next() {  // for循环 遍历rows的值
		err = rows.Scan(&book.Bookid, &book.BookName, &book.CreatTime) //扫描数据
		users = append(users, book)                                    //将book结构体内的数据传入到users切片当中
		if err != nil {
			fmt.Println(err) //错误处理机制
		}
	}
	fmt.Println("aa", users)
	return users
}

// Add 添加书籍
func Add(bookname string, cretime string) int { //通过书籍名称 创建时间参数  //通过server内的Addbook传入bookname cretime
	var a int
	results, err := DB.Exec("insert into book_wyb(bookName,bookTime) values(?,?)", bookname, cretime) //执行 添加sql语句  将bookName bookTime插入表中
	if err != nil {
		panic(err)
	} else { //处理错误机制
		fmt.Println(results)
		a = 1
	}
	return a

}

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
func ChangeBookTime(id int, bookTime string) int { //前端传入的书籍 id 与 书籍时间 作为参数     //通过server内的changeBooktime传入id bookTime
	var a int
	results, err := DB.Exec("update book_wyb set bookTime=? where id =?", bookTime, id) //通过书籍id 修改 书籍时间 执行sql语句 将值传给results
	if err != nil {
		panic(err)
	} else { // 处理错误机制
		fmt.Println(results)
		a = 1
	}
	return a
}
