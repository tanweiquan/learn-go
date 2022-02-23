package main

import "fmt"

//结构体嵌套

type address struct {
	province string
	city     string
}
type workPlace struct {
	province string
	city     string
}
type person struct {
	name string
	age  int
	// province string
	// city     string
	// addr address //这就是结构体的嵌套，把结构体嵌套到另外一个结构体内部
	address   //可以写成匿名嵌套结构体，常用
	workPlace //这里的匿名嵌套结构体跟address里的字段冲突了，但仍可以这样写
}

func main() {
	//结构体嵌套的使用
	p1 := person{
		name: "周玲",
		age:  18,
		/*
			 addr: address{
				province: "广东",
				city:     "广州",
			},
		*/
		address: address{
			province: "广东",
			city:     "广州",
		},
		workPlace: workPlace{
			province: "山东",
			city:     "威海",
		},
	}
	fmt.Println(p1)      //{周玲 18 {广东 广州}}
	fmt.Println(p1.name) //周玲
	// fmt.Println(p1.addr.city) //广州
	//先在结构体内部找变量，找不到就去匿名结构体找变量
	// fmt.Println(p1.city) //匿名字段没有冲突时可以这样写，好处是省略了addr,直接用p1.city就可以拿到city这个变量

	//当有多个匿名嵌套结构体时，且匿名结构体字段冲突(都有city)，就不支持上面这么写了
	//应该写成这样了
	fmt.Println(p1.address.city)

}
