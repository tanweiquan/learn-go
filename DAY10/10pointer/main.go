package main

import "fmt"

//指针也是变量，它也有内存地址：如&x是一个存放x内存地址数的变量，&x也有内存地址
func main() {
	var a = 100
	b := &a                               //声明b是一个存放a内存地址的变量，他的值是一个数
	fmt.Printf("a的类型：%T b的类型：%T\n", a, b) //a的类型：int b的类型：*int
	fmt.Printf("%p\n", &a)                //0xc000014098
	fmt.Printf("%p\n", b)                 //0xc000014098
	//这里打印的是b的值，也就是说b的值是一个数(存放b内存地址的数)
	fmt.Printf("%v\n", b) //0xc000014098
	//这里打印了指针变量的内存地址
	fmt.Printf("%p\n", &b) //0xc000006028
}
