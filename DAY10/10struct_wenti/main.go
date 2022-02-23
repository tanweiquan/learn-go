package main

import "fmt"

type myint int

func (i myint) hello() {
	fmt.Println("hello,这是一个int")
}

type person struct {
	name string
	age  int
}

// //问题3:为什么要有构造函数
// func newPerson(x string, y int) person {
// 	//别人调用，我返回给他一个person类型变量
// 	/* p2 := person{
// 		name: x,
// 		age:  y,
// 	}
// 	return p5 */
// 	return person{
// 		name: x,
// 		age:  y,
// 	}
// }
// func newPerson(x string, y int) *person {
// 	//别人调用，我返回给他一个*person类型变量
// 	return &person{
// 		name: x,
// 		age:  y,
// 	}
// }
func main() {
	//1.myint(100)是个啥？

	//声明1个int32类型的变量，它的值是10
	//方法1：var x int32
	//       x=10
	//方法2：var x int32 = 10
	//方法3：var x=int32(10)
	//方法4：x:=int32(10)
	m := myint(100) //var m myint = 100
	m.hello()
	//问题2：结构体初始化
	//方法1：
	var p1 person
	p1.name = "元帅"
	p1.age = 18
	fmt.Println(p1)
	//方法2
	//键值对初始化
	var p2 = person{
		name: "元帅",
		age:  18,
	}
	fmt.Println(p2)
	// s1 := []int{1, 2, 3}
	/* s2 := map[string]int{
		"stu1": 100,
		"stu2": 200,
		"stu3": 300,
	} */
	//值列表初始化
	var p3 = person{
		"元帅",
		18,
	}
	fmt.Println(p3)
}
