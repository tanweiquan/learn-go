package main

import (
	"fmt"
	"strconv"
)

//strconv库
//数据的类型转换

func main() {
	/* 错误做法：
	f := 97
	strx := int32(f)     //这里将f强转int32类型
	retx := string(strx) //这里变成了通过string()在ASCII表里取值了，取得的是表里对应位置的值a,而不是转化成字符串类型
	fmt.Println(retx)    //a */

	// 把数字转换成字符串类型
	i := 97
	//方法一：
	ret2 := fmt.Sprintf("%d", i)
	fmt.Printf("%v\n", ret2)
	//方法二：(Itoa)
	ret3 := strconv.Itoa(i)
	fmt.Printf("%v\n", ret3)

	// 从字符串中解析出对应的数字
	str := "10000"
	ret1, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("parseint failed,err:", err)
		return
	}
	fmt.Printf("%v %T\n", ret1, int(ret1))

	// 字符串转换成int(Atoi)
	reInt, _ := strconv.Atoi(str)
	fmt.Printf("%v %T\n", reInt, reInt)

	// 从字符串中解析出布尔值
	boolStr := "true"
	boolValue, _ := strconv.ParseBool(boolStr)
	fmt.Printf("%v %T\n", boolValue, boolValue)

	// 从字符串中解析出浮点数
	floatStr := "1.234"
	floatValue, _ := strconv.ParseFloat(floatStr, 64)
	fmt.Printf("%v %T\n", floatValue, floatValue)

}
