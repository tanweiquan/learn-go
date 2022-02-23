package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" //init(),mysql驱动
	"github.com/jmoiron/sqlx"          //第三方的库
)

//区别于Go内置的sql包，第三方的sqlx包简化操作

var db *sqlx.DB //声明db是一个连接池

//设置连接数据库的函数
func initDB() (err error) {
	//数据库信息(数据源)
	dsn := "tanweiquan:qq127468@tcp(127.0.0.1:3306)/sql_test" //前面是用户名和密码，后面的goday14是自己新建的数据库
	//连接数据库
	db, err = sqlx.Connect("mysql", dsn) //Open不会校验用户名和密码是否正确
	if err != nil {                      //dsn的格式不正确的时候才会报错
		return
	}

	db.SetMaxOpenConns(10) //设置连接池的最大连接数
	db.SetMaxIdleConns(5)  //设置最大空闲连接数
	return
}

//先声明一个含有数据表关键字段的结构体
type user struct { //每列都有不同的关键字段
	ID   int //这里要大写，因为外部的sqlx包需要判断结构体里的对象
	Name string
	Age  int
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed,err:%v\n", err)
		return
	}
	fmt.Println("连接数据库成功！")
	//查询单条数据
	var u user
	sqlStr1 := `select id,name,age from user where id=1`
	err = db.Get(&u, sqlStr1)
	if err != nil {
		fmt.Printf("Get failed,err:%v\n", err)
		return
	}
	fmt.Printf("u:%#v\n", u)
	//查询多条数据
	var userList []user
	sqlStr2 := `select id,name,age from user where id>0`
	err = db.Select(&userList, sqlStr2)
	if err != nil {
		fmt.Printf("select failed,err:%v\n", err)
		return
	}
	fmt.Printf("userList:%#v\n", userList)
}
