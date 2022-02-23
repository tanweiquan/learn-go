package main

import "fmt"

//switch
func main() {
	n := 4
	switch n {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default: //默认值
		fmt.Println("无效的数字")
	}
	//在switch里声明变量
	switch m := 3; m {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default: //默认值
		fmt.Println("无效的数字")
	}
	//一次设置多个同种case
	switch x := 5; x {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8, 10:
		fmt.Println("偶数")
	default: //默认值
		fmt.Println(x)
	}
	//在case中使用表达式
	_age := 61
	switch {
	case _age < 25:
		fmt.Println("好好学习吧")
	case _age > 25 && _age < 35:
		fmt.Println("好好工作吧")
	case _age > 60:
		fmt.Println("好好享受吧")
	default: //默认值
		fmt.Println("活着真好")
	}
	switch age := 30; {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default: //默认值
		fmt.Println("活着真好")
	}
}
