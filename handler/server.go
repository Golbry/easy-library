package handler

import (
	"easy-library/common"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Index 前端登入注册页面
func Run() {
	engine := gin.Default()                 //返回默认引擎
	engine.Static("/static", "view/static") //加载静态文件
	engine.LoadHTMLGlob("view/html/*")      //加载view/html/所有的html文件
	engine.Use(cors.Default())              //跨域 插件
	//登入页面
	engine.GET("/index", func(c *gin.Context) {
		//根据文件名渲染
		//加载模板是加载的路径，替换的是文件中的某个变量
		c.HTML(200, "index.html", gin.H{"title": "标签页"})
	}) //触发get请求 加载html文件
	//书籍主页
	engine.GET("/back_stage", func(c *gin.Context) {
		//根据文件名渲染
		//加载模板是加载的路径，替换的是文件中的某个变量
		c.HTML(200, "back_stage.html", gin.H{"title": "书籍查询"})
	}) //触发get请求 加载html文件
	//书籍编辑
	engine.GET("/book_set", func(c *gin.Context) {
		//根据文件名渲染
		//加载模板是加载的路径，替换的是文件中的某个变量
		c.HTML(
			200, "book_set.html", gin.H{"title": "书籍管理"})
	}) //触发get请求 加载html文件
	//添加书籍页面
	engine.GET("/add_books", func(c *gin.Context) {
		//根据文件名渲染
		//加载模板是加载的路径，替换的是文件中的某个变量
		c.HTML(200, "add_books.html", gin.H{"title": "添加图书"})
	}) //触发get请求 加载html文件
	easy_library := engine.Group("easy_library") //创建 请求组
	{
		easy_library.POST("/login", Login)            //POST>>>>登入
		easy_library.POST("/register", Register)      //POST>>>>注册
		easy_library.POST("/show_data", ShowData)     //POST>>>>动态展示
		easy_library.POST("/add_book", AddBook)       //POST>>>>添加书籍
		easy_library.POST("/change_book", ChangeBook) //POST>>>>修改书籍
		easy_library.POST("/query_book", QueryBook)   //POST>>>>查询
		easy_library.POST("/borrow_book", BorrowBook) //POST>>>>借阅图书
		easy_library.POST("/return_book", ReturnBook) //POST>>>>归还图书
	}
	err := engine.Run("127.0.0.1:8082") //启动端口127.0.0.1:8082
	if err != nil {
		return
	} //处理错误机制
}
func toHtml(c *gin.Context, cusErr *common.CusError, data interface{}, exception interface{}) { //渲染html 返回
	c.JSON(200, gin.H{
		"error":     cusErr,
		"data":      data,
		"exception": exception,
	})
}
