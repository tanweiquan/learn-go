package main

import "fmt"

//空接口
/*
type xx interface{

}
*/
//空接口没有必要起名字，通常定义成如此格式：interface{}
//所有的类型都实现了空接口，也就是任意类型的变量都能保存到空接口中
//interface：关键字
//interface{}：空接口类型
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}
func main() {
	m1 := make(map[string]interface{}, 16)
	m1["name"] = "周玲"
	m1["age"] = 26
	m1["married"] = true
	m1["hobby"] = [...]string{"唱", "跳", "rap"}
	fmt.Println(m1) //map[age:26 hobby:[唱 跳 rap] married:true name:周玲]

	show(false) //type:bool value:false
	show(nil)   //type:<nil> value:<nil>
	show(m1)    //type:map[string]interface{} value:map[age:26 hobby:[唱 跳 rap] married:true name:周玲]
}
