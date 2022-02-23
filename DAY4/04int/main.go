package main

import "fmt"

//整型
func main() {
	//十进制
	var a = 101
	fmt.Printf("%d\n", a) //十进制数值
	fmt.Printf("%b\n", a) //把十进制转换为二进制
	fmt.Printf("%o\n", a) //把十进制转换为八进制
	fmt.Printf("%x\n", a) //把十进制转换为十六进制

	//八进制（以0开头，后面跟的数值从0到7.共8个数）
	b := 077
	fmt.Printf("%d\n", b)
	//十六进制（以0x开头，后面的数值从0到f，如0xff）
	c := 0x1234567
	fmt.Printf("%d\n", c)
	//查看变量类型
	fmt.Printf("%T\n", c)
	//声明int8类型的常量
	d := int8(9)          //明确指定int8类型，否则默认为int类型
	fmt.Printf("%T\n", d) //打印声明类型为int8
	fmt.Printf("%d", d)   //打印赋值结果是9
}

//int和uint的区别是，int有符号，uint无符号
//在32位操作系统是int32或uint32;在64位系统就是int64或uint64
/*例 int8是有符号的整型（-128到127）
uint是无符号的整型（0到255）*/
//uintptr无符号整型，用于存放一个指针
/*占位符
%T表示类型，
%v表示任何值（含数字、文字，真假描述等），
%d表示整型数字，
%b表示二进制值，
%o表示八进制值，
%x表示十六进制值。
%s表示字符串
%c表示字符
#%s中的#，会加一个描述符
*/
