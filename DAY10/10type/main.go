package main

import "fmt"

//自定义类型和类型别名
//自定义类型的形式如：type 自定义新类型 基本类型
type myInt int //自定义类型
//类型别名的形式如：type 自定义新类型 = 基本类型
type yourInt = int //类型别名

func main() {
	var a myInt = 10      //声明a是myInt类型,且值为10
	fmt.Println(a)        //10
	fmt.Printf("%T\n", a) //类型为myInt
	var b yourInt = 100
	fmt.Println(b)        //100
	fmt.Printf("%T\n", b) //类型为int，yourInt为int的类型别名
	var c rune = '中'
	fmt.Println(c)        //20013
	fmt.Printf("%T\n", c) //类型为int32,rune是int32的类型别名
}
