package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //init(),mysql驱动
)

//事务(保证两个sql语句共同成功，防止一个成功一个失败）
//例如银行转账业务，要保证扣钱和存钱都共同成功，防止出现一个成功扣款，另外一个却没存钱成功的现象

var db *sql.DB //声明db是一个连接池

//Go语言中使用以下三个方法实现MySQL中的事务操作。
/*
开始事务：func (db *DB) Begin() (*Tx, error)     //db.Begin()  -->返回tx
回滚事务：func (tx *Tx) Rollback() error         //tx.Rollback()
提交事务：func (tx *Tx) Commit() error           //tx.Commit()
*/

//设置连接数据库的函数
func initDB() (err error) {
	//数据库信息(数据源)
	dsn := "tanweiquan:qq127468@tcp(127.0.0.1:3306)/sql_test" //前面是用户名和密码，后面的goday14是自己新建的数据库
	//连接数据库
	db, err = sql.Open("mysql", dsn) //Open不会校验用户名和密码是否正确
	if err != nil {                  //dsn的格式不正确的时候才会报错
		return
	}
	err = db.Ping() //Ping尝试连接数据库，可以校验用户名和密码
	if err != nil { //用户名和密码不正确就会报错
		return
	}
	return
}

//事务操作示例
func transactionDemo() {
	//1、开启事务
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			//回滚
			tx.Rollback() //如果开启事务不成功就回滚
		}
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}
	//2、写sql语句
	sqlStr1 := `update user set age=age-1 where id=?`
	sqlStr2 := `update user set age=age+1 where id=?`

	_, err = tx.Exec(sqlStr1, 1) //id为1的age要-1
	if err != nil {
		//回滚事务
		tx.Rollback() //执行sqlStr1语句不成功就回滚
		fmt.Printf("exec sqlStr1 failed, err:%v\n", err)
		return
	}

	_, err = tx.Exec(sqlStr2, 2) //id为2的age要+1
	if err != nil {
		//回滚事务
		tx.Rollback() //执行sqlStr2语句不成功就回滚
		fmt.Printf("exec sqlStr2 failed, err:%v\n", err)
		return
	}
	//3、提交事务
	err = tx.Commit()
	if err != nil {
		tx.Rollback() //执行提交事务不成功就回滚
		fmt.Printf("commit failed,err:%v\n", err)
		return
	}
	fmt.Println("开启事务成功")
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed,err:%v\n", err)
		return
	}
	fmt.Println("连接数据库成功！")
	transactionDemo()
}
