package main

import "fmt"

//go语言中为了处理非ASCII码类型的字符，定义了新的rune类型
func main() {
	s := "hello沙河"
	n := len(s)    //len求的是ASCII字符的个数或字节长度
	fmt.Println(n) //打印结果是11，英文字符是ASCII字符，沙河是汉字，汉字统计的是汉字字节，所以是5+6=11
	str := "abc"
	for i := 0; i < len(str); i++ { //通过for打印每个字符
		fmt.Printf("str[%d]=%c\n", i, str[i])
	}
	for _, c := range s { //从字符串中拿出具体的字符
		fmt.Printf("%c\n", c) //%c是字符
	}
	//字符串修改，
	s2 := "白萝卜"      //['白''萝''卜']
	s3 := []rune(s2) //把字符串强制转换成rune切片
	s3[0] = '红'
	fmt.Println(string(s3)) //把rune切片强制转换为字符串
	c1 := "红"               //string
	c2 := '红'               //rune(int32);Go语言中字符均使用utf-8编码的unicode文本，所以使用rune
	fmt.Printf("c1:%T c2:%T\n", c1, c2)
	h1 := "H"       //string
	h2 := 'H'       //rune(int32);同样,当ASCII字符不转换成byte也默认使用rune（int32）
	h3 := byte('H') //转换为byte类型，此时ASCII字符使用byte(uint8)
	fmt.Printf("h1:%T h2:%T h3:%T\n", h1, h2, h3)
	//类型转换。整型与浮点型转换；字符串与切片转换
	n1 := 10
	f := float64(n1)
	fmt.Println(f)
	fmt.Printf("%T\n", f)
}
