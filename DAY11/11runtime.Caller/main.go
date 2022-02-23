package main

import (
	"fmt"
	"runtime"
)

//runtime.Caller()这个函数主要是为了查看函数调用栈的问题(即是看调用的文件是什么和在哪行输出，没有则输出false)
func f1(n int) {
	pc, file, line, ok := runtime.Caller(n) //n看套了多少层函数
	if !ok {
		fmt.Println("runtime.Caller() failed")
		return
	}
	runtime.FuncForPC(pc)
	fmt.Println(file)
	fmt.Println(line)
}

func f2(n int) {
	f1(n)
}
func main() {
	f1(1) //file:E:/go/src/Learning/DAY11/11runtime.Caller/main.go;line:24
	f2(2) //file:E:/go/src/Learning/DAY11/11runtime.Caller/main.go;line:25
}
