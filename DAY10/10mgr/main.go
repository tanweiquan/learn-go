package main

import (
	"fmt"
	"os"
)

//学生管理系统

//方法版

//菜单函数
func showMenu() {
	fmt.Println("-----------welcome sms!-----------")
	fmt.Println(`
       1、查看所有学生
       2、新增学生
       3、修改学生
       4、删除学生
       5、退出
   `)
}
func main() {
	smr := studentMgr{
		allStudent: make(map[int64]student, 100),
	}
	for {
		//展示菜单
		showMenu()
		//等待用户输入选项
		fmt.Println("请输入序号：")
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你输入的是：%d\n", choice)
		//按选项执行
		switch choice {
		case 1:
			smr.showStudent()
		case 2:
			smr.addStudent()
		case 3:
			smr.editStudent()
		case 4:
			smr.deleteStudent()
		default:
			os.Exit(1)
		}
	}
}
