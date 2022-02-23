package main

import (
	"fmt"
)

//浮点数
/*go里支持两种浮点型：float32和float64。
float32的浮点数最大范围为3.4e38；
float64的浮点数的最大范围为1.8e308。
常量定义
float32:math.MaxFloat.32。
float64:math.MaxFloat.64。
*/
func main() {
	//math.MaxFloat32//float32最大值
	f1 := 1.23456
	fmt.Printf("%T\n", f1) //默认Go语言中的小数都是float64类型
	f2 := float32(1.23456) //显示声明float32类型
	fmt.Printf("%T", f2)
	//错误：f1 = f2，float32类型的值不能直接赋值给float64类型的变量
}
