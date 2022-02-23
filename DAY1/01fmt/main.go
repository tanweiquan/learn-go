package main

import "fmt"

func main() {
	//打印结果:fmt.Print、fmt.Println、fmt.Printf三个函数
	//fmt.Print();同行打印，也就是打印完后不换行
	fmt.Print("沙河")
	fmt.Print("你好")
	fmt.Println(".......")
	//fmt.Println();换行打印，也就是打印完后换行
	fmt.Println("沙河")
	fmt.Println("你好")
	//fmt.Printf("格式化字符串"，值)
	//%T：查看类型
	//%d：十进制数
	//%b：二进制数
	//%o：八进制数
	//%x：十六进制数，使用a-f
	//%X：十六进制数，使用A-F
	//%c：字符
	//%U：表示为Unicode格式
	//%q：该值对应的单引号括起来的go语法字符字面值，必须时会采用安全的转义表示
	//如fmt.printf("%q\n",65) 打印结果为'A',fmt.printf("%q\n","理想") 打印结果为"理想"
	//%s：字符串
	//%p：指针
	//%e：科学计数法，如-1234.456e+78
	//%E：科学计数法，如-1234.456E+78
	//%f：浮点数,默认宽度，默认精度；n:=12.34 fmt.Printf("%f\n",n)//打印结果是：12.34
	//%9f：宽度9，默认精度；n:=12.34 fmt.Printf("%9f\n",n)//打印结果是：12.340000
	//%.2f：默认宽度，精度2；n:=12.34 fmt.Printf("%.2f\n",n)//打印结果是：12.34
	//%9.2f：宽度9，精度2；n:=12.34 fmt.Printf("%9.2f\n",n)//打印结果是：    12.34//前面空4位
	//%9.：宽度9，精度0；n:=12.34 fmt.Printf("%9.f\n",n)//打印结果是：       12//前面空7位
	//%F：等价于%f
	//%g：自动根据实际情况选择使用%e或%f(以获得更简洁、准确的输出)
	//%G：自动根据实际情况选择使用%E或%F(以获得更简洁、准确的输出)
	//%t：布尔值
	//%%：百分号
	//%v：值的默认格式表示
	//%+v：类似%v，但输出结构体时会添加字段名
	//%#v:值的go语法表示
	//........................................................................
	//获取输入：fmt.Scan、fmt.Scanln、fmt.Scanf三个函数
	//fmt.Scan()：获取用户输入
	var a string                //声明a的类型
	fmt.Scan(&a)                //用户输入a的值
	fmt.Println("用户输入的内容是：", a) //打印用户输入的值
	//fmt.Scanf():获取用户的格式化输入
	var (
		name  string
		age   int
		class string
	)
	fmt.Scanf("%s %d %s\n", &name, &age, &class)
	fmt.Println(name, age, class)
	var (
		b string
		c int
	)
	//fmt.Scanln():获取用户的换行输入
	fmt.Scanln(&a, &b)
	fmt.Println(b, c)
}
