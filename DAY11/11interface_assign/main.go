package main

import "fmt"

/*
只有当有两个或两个以上的具体类型必须以相同的方式进行处理时才需要定义接口，
不要为了接口而写接口，那样只会增加不必要的抽象，导致不必要的运行时的消耗。
*/
//空接口可以储存任意类型的值，所以在go语言中的使用十分广泛
//类型断言：我想知道空接口接收的值是什么？

//类型断言1
func assign1(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if !ok {
		fmt.Println("猜错了")
	} else {
		fmt.Println("传进来的是一个字符串：", str)
	}
}

//类型断言2
func assign2(a interface{}) {
	fmt.Printf("%T\n", a)
	switch t := a.(type) {
	case string:
		fmt.Println("是一个字符串：", t)
	case int:
		fmt.Println("是一个int：", t)
	case int64:
		fmt.Println("是一个int64：", t)
	case bool:
		fmt.Println("是一个bool：", t)
	default:
		fmt.Println("查无此类型")
	}

}
func main() {
	//if类型断言
	assign1(100)
	assign1("你好")
	//switch类型断言
	assign2(100)
	assign2(true)
	assign2(int(64))
}
