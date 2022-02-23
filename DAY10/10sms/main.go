package main

import (
	"fmt"
	"os"
)

//函数版学生管理系统
//写一个系统能够查看\新增学生\删除学生
type student struct {
	id   int64
	name string
}

var (
	allStudent map[int64]*student //变量声明
)

func showAllstudent() {
	//把所有的学生都打印出来
	for k, v := range allStudent {
		fmt.Printf("学号：%d 姓名：%s\n", k, v.name)
	}
}

//newStudent 是student类型的构造函数
func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}
func addStudent() {
	//向allStudent中添加一个新的学生
	//1.创建一个新学生
	//1.1获取用户输入
	var (
		id   int64
		name string
	)
	fmt.Print("请输入学生学号：")
	fmt.Scanln(&id)
	fmt.Print("请输入学生姓名：")
	fmt.Scanln(&name)
	//1.2造学生
	newstu := newStudent(id, name)
	//2.追加到allStudent这个map中
	allStudent[id] = newstu

}
func deleteStudent() {
	//1.请用户输入要删除的学生的序号
	var (
		deleteID int64
	)
	fmt.Print("请输入学生学号：")
	fmt.Scanln(&deleteID)
	//2.去allStudent这个map中根据学号删除对应的键值对
	delete(allStudent, deleteID)
}

func main() {
	allStudent = make(map[int64]*student, 48) //初始化(开辟内存空间)
	for {                                     //用for循环来使程序执行后不退出，继续执行程序
		//1.打印菜单
		fmt.Println("欢迎光临学生管理系统！")
		//2.等待用户下一步要干什么
		fmt.Println(`
	1.查看所有学生
	2.新增学生
	3.删除学生
	4.退出
	`)
		fmt.Print("请输入你要干啥：")
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了%d这个选项\n", choice)

		//3.执行对应的函数
		switch choice {
		case 1:
			showAllstudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1) //退出
		default:
			fmt.Println("输入错误")
		}
	}
}
