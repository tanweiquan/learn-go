package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" //init(),mysql驱动
	"github.com/jmoiron/sqlx"          //第三方的库
)

//SQL注入
//SQL注入问题:我们任何时候都不应该自己拼接SQL语句！
var db *sqlx.DB //声明db是一个连接池

func initDB() (err error) {
	//数据库信息(数据源)
	dsn := "tanweiquan:qq127468@tcp(127.0.0.1:3306)/sql_test" //前面是用户名和密码，后面的goday14是自己新建的数据库
	//连接数据库
	db, err = sqlx.Open("mysql", dsn) //Open不会校验用户名和密码是否正确
	if err != nil {                   //dsn的格式不正确的时候才会报错
		return
	}
	err = db.Ping() //Ping尝试连接数据库，可以校验用户名和密码
	if err != nil { //用户名和密码不正确就会报错
		return
	}
	return
}

type user struct { //每列都有不同的关键字段
	ID   int
	Name string
	Age  int
}

//SQL注入示例
func sqlInjectDemo(name string) {
	//模拟拼接的sql语句
	//真正写代码不要这样写
	sqlStr := fmt.Sprintf("select id,name,age from user where name='%s'", name)
	//打印出拼接的SQL语句
	fmt.Printf("SQL:%s\n", sqlStr)
	var userList []user
	err := db.Select(&userList, sqlStr)
	if err != nil {
		fmt.Printf("exec failed,err:%v\n", err)
		return
	}

	for _, v := range userList {
		fmt.Printf("userList%v:%v\n", v, userList)
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed,err:%v\n", err)
		return
	}
	fmt.Println("连接数据库成功！")
	sqlInjectDemo("房飞凤")
	//SQL注入问题:我们任何时候都不应该自己拼接SQL语句！
	//解决方法是做预处理，使用sql包里的Prepare()方法
	/*SQL:select id,name,age from user where name='xxx'or 1=1#';or 1=1必成立，而且有#注释掉'号，导致输入的SQL语句变成了select id,name,age from user
	sqlInjectDemo("xxx'or 1=1#") */

	/*
		此时以下输入字符串都可以引发SQL注入问题：
		sqlInjectDemo("xxx' or 1=1#")
		sqlInjectDemo("xxx' union select * from user #")
		sqlInjectDemo("xxx' and (select count(*) from user) <10 #")
	*/
}
