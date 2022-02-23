package main

import (
	"fmt"
	"strings"
)

//闭包进阶

func f1(f func()) { //参数类型为func()
	f()
}
func f2(x, y int) { //函数类型为func(int,int)
	fmt.Println(x + y)
}

//要求f2作为参数传进f1中，f1(f2)
//为了把函数f2作为参数传进f1中,因此可以定义下面的函数
func f3(f func(int, int), x, y int) func() {
	tmp := func() {
		//fmt.Println(x + y) //这里就是f2()中的方法
		f(x, y) //这里当f2传进来时，执行f2(x,y)，即调用了自身
	}
	return tmp //返回值类型为func(),现在可以作为参数传到f1中啦
}

//例子1
//关键字 函数名 (参数名 参数类型)(返回值 返回值类型){}
func f4(x int) (func(int) int, func(int) int) {
	add := func(i int) int {
		x += i
		return x
	}
	sub := func(i int) int {
		x -= i
		return x
	}
	return add, sub
}

//例子2
func f5(suffix string) func(string) string { //该函数用于判断name是否有后缀名，有就直接返回，没有就补全再返回
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
func main() {
	//f3(f2, 200, 300)，f3中传入函数f2，200,300作为参数
	ret := f3(f2, 200, 300) //把原来需要传递两个int类型的参数包装到一个不需要传参的函数
	ret()                   //500
	fmt.Printf("%T\n", ret) //类型为func()符合f1中的参数类型要求
	f1(ret)                 //500
	//定义闭包，a,b为一个闭包，因为需要给两个匿名函数传参数，因此需要两个变量一起去作为一个闭包
	a, b := f4(10)              //传参给了f4中的x,x=10
	fmt.Printf("%T,%T\n", a, b) //类型为func(int) int,func(int) int
	fmt.Println(a(2), b(3))     //12,9；10+2=12，12-3=9，即先执行完a(2)后，再执行b(3)
	fmt.Println(a(3), b(4))     //12,8；上面执行完后的返回值会储存并影响到这里，所以有9+3=12，12-4=8
	fmt.Println(a(3), b(1))     //11,10;同上面，有8+3=11，11-1=10
	s1 := f5(".txt")            //把".txt"作为参数传入f5中，作为suffix的值
	fmt.Println(s1("呵呵哒"))      //呵呵哒.txt
	fmt.Println(s1("哈哈.txt"))   //哈哈.txt
}
