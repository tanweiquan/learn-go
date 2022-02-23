package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//GO连接MYSQL示例

//连接池是将已经创建好的连接保存在池中。
//当有请求来时，直接使用已经创建好的连接对数据库进行访问。这样省略了创建连接和销毁连接的过程。
var db *sql.DB //声明db是一个连接池
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

	/* 设置与数据库建立连接的最大数目。 如果n大于0且小于最大闲置连接数，会将最大闲置连接数减小到匹配最大开启连接数的限制。
	如果n<=0，不会限制最大开启连接数，默认为0（无限制）。*/
	db.SetMaxOpenConns(0) //设置最大连接数

	/* SetMaxIdleConns设置连接池中的最大闲置连接数。(没有任何请求时在连接池中可以存在的连接数)
	如果n大于最大开启连接数，则新的最大闲置连接数会减小到匹配最大开启连接数的限制。如果n<=0，不会保留闲置连接。*/
	db.SetMaxIdleConns(8) //设置最大空闲连接数
	return
}

//先声明一个含有数据表关键字段的结构体
type user struct { //每列都有不同的关键字段
	id   int
	name string
	age  int
}

//查询

//查询单条数据
func queryOne(id int) {
	//声明一个变量是一个结构体，用来接收数据库返回的数据
	var str user

	//1.写查询单条记录的sql语句
	sqlStr := `select id,name,age from user where id=?;`

	/*单行查询db.QueryRow()执行一次查询，并期望返回最多一行结果（即Row）。QueryRow总是返回非nil的值，
	  直到返回值的Scan方法被调用时，才会返回被延迟的错误。因此用完ROW后必须用Scan释放数据库连接（如：未找到结果），
	  注意这里是释放数据库的连接，但没有关闭数据库*/
	/*
		//2.1执行
		   rowObj := db.QueryRow(sqlstr, id) //从连接池里拿一个连接出来去数据库查询单条记录

		   //2.2拿到结果
		   rowObj.Scan(&str.id, &str.name, &str.age)//必须对rowObj对象调用Scan()方法，因为该方法会释放数据库连接
	*/

	// 2.执行并拿到结果
	//上面可以简写成：
	db.QueryRow(sqlStr, id).Scan(&str.id, &str.name, &str.age)

	//3.打印结果
	fmt.Printf("str:%v\n", str)
}

//查询多条数据
func queryMore(id int) {
	//1.写查询多条记录的sql语句
	sqlStr := `select id,name,age from user where id>?;`
	//2.执行
	rows, err := db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("query failed,err:%v\n", err)
		return
	}
	//3.一定要关闭rows,释放数据库连接
	defer rows.Close()
	//4.循环取值
	for rows.Next() {
		var str user
		err := rows.Scan(&str.id, &str.name, &str.age)
		if err != nil {
			fmt.Printf("Scan failed,err:%v\n", err)
			break
		}
		fmt.Printf("str:%v\n", str)
	}

}

//插入
func insert(newName string, newAge int) {
	//1.写sql语句
	sqlStr := `insert into user(name,age) values(?,?)`
	//2.执行exec
	ret, err := db.Exec(sqlStr, newName, newAge)
	if err != nil {
		fmt.Printf("exec failed,err:%v\n", err)
		return
	}
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return
	}
	fmt.Printf("id:%d\n", id)
}

//更新(更改)
func updateRow(newAge int, id int) {
	//1.写sql语句
	sqlStr := `update user set age=? where id=?`
	//2.执行
	ret, err := db.Exec(sqlStr, newAge, id)
	if err != nil {
		fmt.Printf("exec failed,err:%v\n", err)
		return
	}
	//3.查看结果
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return
	}
	fmt.Printf("更新了%d行数据\n", n)
}

//删除
func deleteRow(id int) {
	//1.写sql语句
	sqlStr := `delete from user where id=?`
	//2.执行
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete id failed,err:%v\n", err)
		return
	}
	//3.查看结果
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return
	}
	fmt.Printf("删除了%d行数据\n", n)

}

//预处理
/*
1、优化MySQL服务器重复执行SQL的方法，可以提升服务器性能，提前让服务器编译，一次编译多次执行，节省后续编译的成本。
2、避免SQL注入问题。
*/
//Prepare方法会先将sql语句发送给MySQL服务端，返回一个准备好的状态用于之后的查询和命令。返回值可以同时执行多个查询和命令。
// 预处理查询示例
func prepareQueryDemo(id int) {
	sqlStr := `select id, name, age from user where id > ?`
	stmt, err := db.Prepare(sqlStr) //先把命令发给mysql预处理一下
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(id) //再把数据发给mysql，返回含查询的数据
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

// 预处理插入示例(插入、更新和删除操作的预处理十分类似，这里以插入操作的预处理为例：)
func prepareInsertDemo() {
	sqlStr := "insert into user(name, age) values (?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()

	/* 插入单条数据
	_, err = stmt.Exec("肖源方",26)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	} */

	//插入多条数据
	var m = map[string]int{ //当想在函数外面传值时，可把m作为参数，在传参时声明是map类型
		"刘柒仟": 28,
		"黄忠者": 23,
		"谭嘉卉": 30,
	}
	for k, v := range m {
		stmt.Exec(k, v)
	}

	fmt.Println("insert success.")
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed,err:%v\n", err)
		return
	}
	fmt.Println("连接数据库成功！")

	queryOne(1)  //传id，拿到id为1的那一行记录
	queryMore(0) //传id，拿到id>0的那几行记录
	insert("房飞凤", 25)
	updateRow(26, 2)
	deleteRow(4)
	prepareQueryDemo(1)
	prepareInsertDemo()
}
