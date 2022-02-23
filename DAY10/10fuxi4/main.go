package main

import (
	"encoding/json"
	"fmt"
)

//结构体嵌套
type addr struct {
	province string
	city     string
}
type student struct {
	name string
	addr //匿名嵌套别的结构体，就使用类型做名称
}

func main() {
	s1 := student{
		name: "小明",
		addr: addr{
			province: "广东",
			city:     "广州",
		},
	}
	fmt.Println(s1) //{小明 {广东 广州}}

	//json序列化与反序列化
	//出现问题：
	//1.结构体内的字段要大写!!!不大写别人是访问不到的(给字段改字段名：`json:"自定义字段名"`)
	//2.反序列化时要传递指针

	//序列化：
	type point struct {
		X int `json:"x"`
		Y int `json:"y"`
	}
	p1 := point{100, 200}
	b, err := json.Marshal(p1) //go语言中把错误当成值返回，错误通常作为第二个返回值

	if err != nil { //如果错误不等于空，也就是发生了错误，则执行：
		fmt.Printf("marshal failed,err:%v\n", err) //打印错误
	}
	fmt.Println(string(b)) //{"x":100,"y":200}

	//反序列化：
	str := `{"X":10,"Y":20}`                 //json格式的字符串
	var p2 point                             //造一个结构体变量，准备接收反序列化返回的值
	err1 := json.Unmarshal([]byte(str), &p2) //使用[]byte，将字符串类型转化成字节类型切片
	if err1 != nil {
		fmt.Printf("marshal failed,err1:%v\n", err1) //打印错误
	}
	fmt.Println(p2) //{10 20}

}
