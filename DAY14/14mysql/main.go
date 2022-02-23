package main

//先在终端中下载依赖：go get -u github.com/go-sql-driver/mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //导入包但不直接使用包；使用了这个包里的init(),注册使用mysql
)

//GO连接MYSQL示例

func main() {
	// DSN:Data Source Name
	//数据库信息(数据源)
	dsn := "tanweiquan:qq127468@tcp(127.0.0.1:3306)/goday14" //前面是用户名和密码，后面的goday14是自己新建的数据库
	//连接数据库
	db, err := sql.Open("mysql", dsn) //Open不会校验用户名和密码是否正确
	if err != nil {                   //dsn的格式不正确的时候才会报错
		fmt.Printf("dsn格式不正确，err:%v\n", err)
		return
	}
	defer db.Close() // 注意这行代码要写在上面err判断的下面，因为万一err不为空，则db为空，空的db没有Close()方法
	err = db.Ping()  //Ping尝试连接数据库，可以校验用户名和密码
	if err != nil {  //用户名和密码不正确就会报错
		fmt.Printf("dsn用户名或密码不正确，err:%v\n", err)
		return
	}
	fmt.Println("连接数据库成功！")
	//接下来就可以用返回的db干活了
	//.........
}
