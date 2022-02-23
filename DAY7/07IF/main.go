package main

import "fmt"

//条件语句 if else
func main() {
	age := 19
	if age > 18 {
		fmt.Println("你成年啦！") //if语句中对age判断，符合age>18语句则执行if{}里的代码
	} else {
		fmt.Println("快写作业去！") //不符合age>18语句则执行else{}里的代码
	} //多个判断条件时
	/*格式为
	  if 判断表达式1的布尔值是true{
	  	分支1
	  }else if 判断表达式2的布尔值是true{
	  	分支2
	  }
	  else {
	  	分支3
	  }
	*/
	if age > 35 {
		fmt.Println("人到中年")
	} else if age > 18 {
		fmt.Println("青年人")
	} else {
		fmt.Println("快写作业")
	}
	//作用域，变量a此时只在if条件判断语句中生效，这样做的好处是减少内存占用
	if a := 19; a > 18 {
		fmt.Println("a>18")
	} else {
		fmt.Println("a<18")
	}
	//fmt.Println(a) 这里就找不到变量a了
}
