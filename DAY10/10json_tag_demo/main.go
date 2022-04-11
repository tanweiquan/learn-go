package main

import (
	"fmt"
	"reflect"
)

// tag设置标签/属性
// x  type `key01:"value01" key02:"value02" key03:"value03"`
// 序列化时，程序自动将field(字段)转化为key属性下的value格式的字段
// 反序列化时，程序自动将key属性下的value格式的字段转换为field(字段)

// 有时候会在字段定义后面带上一个字符串(tag),给字段指定属性。
/*  三种获取 field
field := reflect.TypeOf(obj).FieldByName("Name")
field := reflect.ValueOf(obj).Type().Field(i)  // i 表示第几个字段
field := reflect.ValueOf(&obj).Elem().Type().Field(i)  // i 表示第几个字段

// 获取 Tag
tag := field.Tag

// 获取键值对
labelValue := tag.Get("label")
labelValue,ok := tag.Lookup("label") */
type Person struct {
	Name   string `key:"Name is: "`
	Age    int    `key:"Age is: "`
	Gender string `key:"Gender is: " default:"unknown"`
}

func main() {
	person := Person{
		Name: "MING",
		Age:  29,
	}
	// 取 Value
	v := reflect.ValueOf(person)
	// 解析字段
	for i := 0; i < v.NumField(); i++ {

		// 取tag
		field := v.Type().Field(i)
		tag := field.Tag

		// 解析key 和 default
		label := tag.Get("key")
		defaultValue := tag.Get("default")

		value := fmt.Sprintf("%v", v.Field(i))
		if value == "" {
			// 如果没有指定值，则用默认值替代
			value = defaultValue
		}
		fmt.Println(label + value)
	}
}
