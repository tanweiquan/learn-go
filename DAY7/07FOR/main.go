package main

import "fmt"

//for（循环结构）
//语句 1 在循环（代码块）开始前执行 (i:=?)
//语句 2 定义运行循环（代码块）的条件 (i<?)
//语句 3 在循环（代码块）已被执行之后执行 (i++)
func main() {
	//基本格式
	for i := 0; i < 10; i++ { //for循环中对i判断,若符合i<10语句则执行for{}里的,当执行完{}里的代码后再执行i++，然后再判断符合i<10，符合则继续循环
		fmt.Println(i)
	}
	//变种1
	var g = 5
	for ; g < 10; g++ {
		fmt.Println(g)
	}
	//变种2
	var h = 11
	for h < 15 {
		fmt.Println(h)
		h++
	}
	/*
		无限死循环：(不要尝试)
		for {
			fmt.Println("123")
		}
	*/
	//for range(键值循环)遍历数组，切片，字符串,map及通道（channel）。根据目录找内容
	s := "hello沙河"
	//1.数组，切片，字符串返回索引和值；2.map返回键和值；3.通道（channel）只返回通道内的值
	for m, n := range s {
		fmt.Printf("%d %c\n", m, n)
	}

	//9x9乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%dx%d=%d\t", j, i, i*j) //其中\t是水平制表符
		}
		fmt.Println()
	}
	for i := 1; i < 10; i++ {
		for j := 9; j >= i; j-- {
			fmt.Printf("%dx%d=%d\t", i, j, i*j) //其中\t是水平制表符
		}
		fmt.Println()
	}
	//流程控制之跳出循环
	//使用break跳出循环，例如当i=5时跳出for循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			break //跳出循环
		}
		fmt.Println(i) //打印0到4,也就是说，当i=5时，跳出for循环里的打印语句
	}
	fmt.Println("over")
	//当i=5时跳过此次循环，不执行for循环内的打印语句，继续下次循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue //继续下次循环
		}
		fmt.Println(i) //跳过打印i=5，继续打印下面的值
	}
	fmt.Println("over")
	//跳出多层for循环
	var flag = false //用flag接收了布尔值，且值为false
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				flag = true //重新给flag赋值为true
				break       //跳出内层for循环
			}
			fmt.Printf("%v-%c\n", i, j)
		}
		//一个布尔类型的值只有两种：true 或 false。
		//if 和 for 语句的条件部分都是布尔类型的值，并且==和<等比较操作也会产生布尔型的值。
		//实际开发中我们应尽量采用比较简洁的布尔表达式，就像用 x 来表示x==true。
		if flag { //判断表达式中flag是布尔值为true,即表达式正确，执行下面式子
			break //跳出for循环（外层的for循环）
		}
	}
	fmt.Println("over")
	//goto+label实现跳出多层for循环
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				goto xx //跳到指定的那个label标签
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}
xx: //label标签
	fmt.Println("over")
}
