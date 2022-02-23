package main

import "fmt"

//go语言中函数的return不是原子操作，在底层是分为两步来执行的(注意这里看return,不是看整个函数)
//第一步：返回值赋值
//defer
//第二步：真正的return返回值
//注意重点：返回值赋值后不能再被赋值，但返回值赋值后可以通过defer来增值或减值
func f1() int { //假设让未命名的返回值统一称为M
	x := 5
	defer func() {
		x++
	}()
	return x //返回值未命名的：首先函数内执行x:=5,然后执行return x(返回值M赋值M=x=5;然后defer里x++,最后返回值M=5)
}
func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 //返回值已经命名的：首先执行return 5（返回值x赋值x=5,然后defer里x++，最后返回值x=6)
}
func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x //返回值已命名：首先函数内执行了x:=5;然后执行return x(返回值y赋值y=x=5,然后defer里x++,最后返回值y=5)
}
func f4() (y int) {
	x := 5
	defer func() {
		y++
	}()
	return x //返回值已命名：首先函数内执行了x:=5;然后执行return x(返回值y赋值y=x=5,然后defer里y++,最后返回值y=6)
}
func f5() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5 //返回值已命名：首先函数内执行return 5(返回值x赋值x=5,然后defer，最后返回值x=5)
}

func main() {
	fmt.Println(f1()) //5
	fmt.Println(f2()) //6
	fmt.Println(f3()) //5
	fmt.Println(f4()) //6
	fmt.Println(f5()) //5
}
