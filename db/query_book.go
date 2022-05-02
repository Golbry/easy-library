package db

func QueryBook(id int) []BookInfo { //通过id查询  返回 books切片  通过server内的Querybook传入id
	//users := make([]books_wyb, 0)                                    //创建 users空切片
	//rows, err := DB.Query("SELECT  *  from book_wyb where id=?", id) //通过id执行查询 sql语句
	//if err != nil {
	//	panic(err)
	//} //执行错误机制
	//var book books_wyb //将结构体 定义为 book
	//for rows.Next() {  // for循环 遍历rows的值
	//	err = rows.Scan(&book.Bookid, &book.BookName, &book.CreatTime) //扫描数据
	//	users = append(users, book)                                    //将book结构体内的数据传入到users切片当中
	//	if err != nil {
	//		fmt.Println(err) //错误处理机制
	//	}
	//}
	//fmt.Println("aa", users)
	//return users
	return nil
}
