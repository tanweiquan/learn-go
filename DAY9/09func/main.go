package main

import "fmt"

//go语言支持函数，匿名函数和闭包
/*
定义函数的格式（即自己创造函数）：
func 函数名 (参数变量，参数变量类型)(返回值变量，返回值变量类型){
      函数体
}
简化记忆：
func 函数名(参数)(返回值){
	函数体
}
多个参数用逗号","分隔;
多个返回值用()包裹，并用逗号","分隔。
返回值可以只写返回值类型，不写变量名称。
Go语言中通过return关键字向外输出返回值。
*/
/*例子：
func sum(x int, y int) (z int) { //定义一个sum(x,y)的函数,其中的x,y都是int，其中返回值变量z可以不写，只写返回值的类型
	return x + y
}
func main() {
	r := sum(1, 2) //声明r为使用sum函数的变量
	fmt.Print(r)   //3，打印结果为函数返回的return值:x+y
}
*/

//没有返回值
func f1(x int, y int) {
	fmt.Println(x + y) //当自定义的函数没有设置return来返回值时，后面可以直接使用函数
}

//没有参数，没有返回值
func f2() {
	fmt.Println("f2的值") //当自定义的函数没有设置return来返回值时，后面可以直接使用函数
}

//没有参数，但有返回值
func f3() int {
	h := 3
	return h
}

//返回值可以命名，也可以不命名
func f4(x int, y int) (r int) {
	r = x + y
	return
}
func f5(x int, y int) int {
	return x + y
}

//返回值命名后，先赋值后return
func f6(x int, y int) (p int) {
	p = x + y
	return p
}

//多个返回值
func f7() (string, int) {
	return "沙河", 1 //这里有两个返回值
}

//参数类型的简写(当多个参数变量的类型一致时，可以将非末尾的参数类型省略)
func f8(a, b, c int, m, n string, i, j bool) (int, int, int) { //注意下面return有多少个值就声明多少个返回值类型
	return a, b, c
}

//可变长参数
//可变长参数必须放在函数参数的最后
func f9(x string, y ...int) {
	fmt.Print(x)
	fmt.Println(y) //y类型是切片[]int
}

//go语言中没有默认参数这个概念
func main() {
	//函数的调用，格式为：函数名()，例如：函数名(元素x,元素y)调用后不一定会直接输出值
	//a1 := f1(2, 3) //没有返回值，因此该自定义函数不能被用来声明一个变量a1
	f1(2, 3) //5,直接调用函数就好，因为已经用了打印方法打印值了
	//a2 := f2() //没有参数，要声明的变量a2也没有参数变量，没有返回值，因此该自定义函数用来声明变量
	f2()                                                 //直接调用函数就好，不需要fmt包和给返回值进行输出
	a3 := f3()                                           //定义的函数没有参数，所以要调用函数的变量a3也应该是没有参数变量；有返回值，可用于声明变量
	fmt.Println(a3)                                      //3,打印返回值为h的值
	a4 := f4(2, 6)                                       //有返回值，按正常声明变量
	fmt.Println(a4)                                      //8,返回值为r的值
	a5 := f5(5, 6)                                       //有返回值，按正常声明变量
	fmt.Println(a5)                                      //11,返回值为x+y
	a6 := f6(4, 9)                                       //有返回值，正常变量声明
	fmt.Println(a6)                                      //13,返回值为p,也就是x+y
	a7, a8 := f7()                                       //有两个返回值，所以接收方也需要两个变量接收
	fmt.Println(a7, a8)                                  //打印结果为 沙河 1
	a9, a10, a11 := f8(1, 2, 3, "你好", "嗯嗯", true, false) //有返回值，正常声明变量
	fmt.Println(a9, a10, a11)                            //打印结果为1，2，3
	f9("下雨啦")                                            //下雨啦[]
	f9("下雨啦", 1, 2, 4, 8)                                //下雨啦[1，2，4，8]
}
