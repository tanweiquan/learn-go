package main

import "fmt"

//结构体是值类型
type person struct {
	name, gender string
}

//go语言中函数传参永远是拷贝
func f(x person) {
	x.gender = "女" //修改的的是拷贝的副本的gender
}

//通过内存地址来修改值
func f2(x *person) {
	//(*x).gender = "女" //传进内存地址后转为取值，然后修改值为女(也就是将内存地址重新指向到新值)
	x.gender = "女" //语法糖(便捷写法），go语言自动帮我们转换为上面的格式
}

func main() {
	//var p person
	//p.name = "周林"
	//p.gender = "男"
	p:=person{
		name:"周林",
		gender: "男",
	}

	f(p)                  //这里改不成女是因为这里改的是副本
	fmt.Println(p.gender) //男
	f2(&p)                //这里传进内存地址
	fmt.Println(p.gender) //女

	//结构体指针1
	//new开辟内存空间，返回的是指针，指针是一个变量，它的值是一个内存地址，同时它也有自己的内存地址
	var p2 = new(person) //这里声明了p2是一个指针
	(*p2).name = "理想"
	p2.gender = "保密"       //语法糖，go自动帮我们转换为上面的格式，常使用
	fmt.Printf("%T\n", p2) //*main.person
	//这里打印的是指针p2的值，也就是p2存放了一个内存地址的数
	fmt.Printf("%p\n", p2) //0xc000020400
	//这里打印的是p2的内存地址
	fmt.Printf("%p\n", &p2) //0xc000006030

	//结构体指针2
	//2.1key-value初始化
	var p3 = person{ //如果要拿到person的指针，可以在person前加&，如：var p3 =&person{}
		name:   "元帅",
		gender: "男",
	}
	fmt.Printf("%#v\n", p3) //main.person{name:"元帅", gender:"男"}
	//2.2使用值列表的形式初始化
	//值的顺序要和结构体定义时字段的顺序一致
	p4 := person{
		"陈凡",
		"男",
	}
	fmt.Printf("%#v\n", p4) //main.person{name:"陈凡", gender:"男"}
}
