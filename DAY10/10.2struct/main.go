package main

import "fmt"

//结构体占用一块连续的内存空间
type x struct {
	a int8 //8bit->1byte
	b int8
	c int8
}
type y struct {
	d int8
	f int8
	g int8
	h string //扩展知识：结构体中go语言会进行内存对齐
}

func main() {
	m := x{
		a: int8(10),
		b: int8(20),
		c: int8(30),
	}
	//打印他们结构体里的字段的值，证明了他们的内存地址时连续的
	fmt.Printf("%p\n", &(m.a)) //0xc000014098
	fmt.Printf("%p\n", &(m.b)) //0xc000014099
	fmt.Printf("%p\n", &(m.c)) //0xc00001409a
	n := y{
		d: 40,
		f: 60,
		g: 90,
		h: "哈哈",
	}
	fmt.Printf("%p\n", &(n.d)) //0xc000004078
	fmt.Printf("%p\n", &(n.f)) //0xc000004079
	fmt.Printf("%p\n", &(n.g)) //0xc00000407a
	fmt.Printf("%p\n", &(n.h)) //0xc000004080//扩展知识：go语言对结构体进行了内存对齐
}
