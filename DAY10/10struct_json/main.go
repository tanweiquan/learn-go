package main

import (
	"encoding/json"
	"fmt"
)

//结构体与json
//1.序列化：把go语言中的结构体变量-->json格式的字符串  (从而能让其他语言识别，如JS)
//2.反序列化：把json格式的字符串-->go语言中能够识别的结构体变量

//首先要定义一个结构体
type person struct {
	Name string
	Age  int
}

//..................这里举例如何让输出的字符串首字母变小写
type dog struct {
	Name string `json:"name"` //加`json:"name"`这个，就可以让输出的json格式中的Name变成name
	Age  int    `json:"age"`  //加`json:"age"`这个，就可以让输出的json格式中的Age变成age
}

func main() {
	//然后要造一个结构体类型的变量
	p1 := person{
		Name: "周玲",
		Age:  18,
	}
	//现在想把p1这个结构体转换为一个字符串
	//序列化用json.Marshal()
	b, err := json.Marshal(p1) //序列化//这里要把p1这个变量取到json包里,所以字段名要大写
	if err != nil {
		fmt.Printf("marshal failed，err:%v", err)
		return
	}
	fmt.Printf("%#v\n", string(b)) //"{\"Name\":\"周玲\",\"Age\":18}"
	fmt.Println(string(b))         //{"Name":"周玲","Age":18}//这种就是json格式的字符串了

	//反序列化
	str := `{"name":"理想","age":18}` //这里假设存在一个json格式的字符串
	var p2 person
	_ = json.Unmarshal([]byte(str), &p2) //传指针是为了在函数json.Unmarshal()内部修改p2的值
	fmt.Printf("%#v\n", p2)              //main.person{Name:"理想", Age:18}
	fmt.Println(p2)                      //{理想 18}
	//...............................
	//...............................
	d1 := dog{
		Name: "小黄",
		Age:  5,
	}
	b1, err1 := json.Marshal(d1)
	if err1 != nil {
		fmt.Printf("marshal failed，err1:%v", err1)
		return
	}
	fmt.Printf("%#v\n", string(b1)) //"{\"name\":\"小黄\",\"age\":5}"
	fmt.Println(string(b1))         //{"name":"小黄","age":5
}
