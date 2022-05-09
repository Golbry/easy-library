# easy-library
DUT大三数据库大作业，采用Gin+database/sql包完成
## 环境说明
go version>=1.16
## 运行说明
1. git clone后cd进入项目
2. 在本地创建easy-library数据库及同名数据表，按db/ddl目录下sql文件定义关系模式
3. 输入go mod tidy导入相关包
4. 输入go run main.go开始运行，运行后在浏览器输入网址http:/localhost:8082/index即可开始访问
