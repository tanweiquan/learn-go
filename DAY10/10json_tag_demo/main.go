package main

import (
	"fmt"
	"reflect"
)

// tag是设置标签
// 有时候会在字段定义后面带上一个字符串(tag)。类似于如下

type T struct {
	f1     string "f one"
	f2     string
	f3     string `f three`
	f4, f5 int64  `f four and five` // field定义时候两个名字公用一个属性
}

// Tag在运行时可以通过反射:reflection包来读取
func main() {
	t := reflect.TypeOf(T{})
	f1, _ := t.FieldByName("f1")
	fmt.Println(f1.Tag) // f one
	f2, _ := t.FieldByName("f2")
	fmt.Println(f2.Tag) //
	f4, _ := t.FieldByName("f4")
	fmt.Println(f4.Tag) // f four and five
	f5, _ := t.FieldByName("f5")
	fmt.Println(f5.Tag) // f four and five
}
