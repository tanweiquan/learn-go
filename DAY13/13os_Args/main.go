package main

import (
	"fmt"
	"os"
)

//读取命令行参数
//go run main.go a b c 执行结果是第一个元素是这个包的路径
//13os_Args.exe a b c
//os.Atgs切片的第一个元素是可执行的文件的文件名，后面的元素是命令行程序员自己追加的元素
func main() {

	//os.Args是切片
	fmt.Println(os.Args[0], os.Args[2])
	fmt.Println(os.Args)
	fmt.Printf("%v\n", os.Args)
	fmt.Printf("%T\n", os.Args)
}
