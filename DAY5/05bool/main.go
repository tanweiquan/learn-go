package main

import "fmt"

//布尔值
func main() {
	a := true
	var b bool                        //默认是false
	fmt.Printf("%T\n", a)             //bool
	fmt.Printf("%T value:%v\n", a, b) //ture,value:false
}
