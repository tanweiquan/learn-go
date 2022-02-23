package main

import "fmt"

/* type person struct {
	name string
	age  int
} */

//结构体匿名字段
//可以用于字段比较少也比较简单的场景
//不常用！！！
type person struct {
	string
	int
}

func main() {
	p1 := person{
		"周琳",
		18,
	}
	fmt.Println(p1)
	fmt.Println(p1.string) //这里假如结构体里还有另外的string类型的数据，那就会报错，因此匿名字段不常用
	fmt.Println(p1.int)
}
