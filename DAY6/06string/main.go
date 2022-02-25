package main

import (
	"fmt"
	"strings"
)

func main() {
	//字符串，双引号包裹
	a := "hello,quan"
	fmt.Println(a)
	//字符，单引号包裹
	b := 'w'
	//当我们直接输出字符的值的时候，就是输出了对应字符的码值
	fmt.Println(b)
	//当我们希望出书对应字符时，需要使用格式化输出,%c是字符的占位符
	fmt.Printf("%c\n", b)
	//\本来就有特殊含义的，我应该告诉程序我写的\就是单纯的\
	path := "\"F:\\WeChat\\locales\""
	fmt.Println(path)
	c := "I'm ok"
	fmt.Println(c)
	//多行的字符串
	//用`定义字符串可以防止攻击，写什么打印什么
	d := `
	大家好，
	      我是渣渣灰
		  fmt.Print("haha")
		  "F:\WeChat\locales"
	`
	fmt.Println(d)
	//字符串相关操作
	//求长度
	fmt.Println(len(a))
	//直接拼接字符串 +或fmt.Sprintf
	name := "理想"
	word := "dsb"
	ss := name + word
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s", name, word) //ss1接受fmt.Sprintf函数完成后的返回值
	fmt.Printf("%s%s\n", name, word)
	fmt.Println(ss1)
	//分割
	f := strings.Split(path, "\\") //用\分割了"F:\WeChat\locales"
	fmt.Println(f)
	//包含
	fmt.Println(strings.Contains(ss, "理性")) //判断ss里有没包含理性,结果是flase
	fmt.Println(strings.Contains(ss, "理想")) //判断ss里有没包含理想,结果是true
	//前缀
	fmt.Println(strings.HasPrefix(ss, "理想")) //判断ss的理想是不是前缀，结果是true
	//后缀
	fmt.Println(strings.HasSuffix(ss, "理想")) //判断ss的理想是不是后缀,结果是flase
	//判断子串出现的位置，位置是从0开始
	ss2 := "abcdeb"
	fmt.Println(strings.Index(ss2, "c"))
	fmt.Println(strings.LastIndex(ss2, "b")) //子串最后出现的位置
	//拼接
	fmt.Println(strings.Join(f, "+")) //用+拼接起来,变成了"F:+WeChat+locales"
	//len计算字符串字节数
	n := "hello,流沙河"
	//求得汉字字节长度，5+1+9,字节的数量是：英文5个，英文符号1个，中文字符9个，共15个
	n1 := len(n) //求字符串n的长度，并把长保存到n1中
	fmt.Println(n1)
	fmt.Printf("%T\n", n1)
	m := "hello，流沙河"
	//求得汉字字节长度，5+3+9,求得的是字节数量：英文5个ASCII字符，中文符号字节3个，中文字节9个，共17个
	m1 := len(m)
	fmt.Println(m1)
	fmt.Printf("%T\n", m1)
	h := "hello,流沙河"
	h1 := len([]rune(h)) //编译正确，求得的是9个字符
	fmt.Printf("%T\n", h1)
	fmt.Println(h1)
	h2 := len([]byte(h)) //全求的是字节数量，5+1+9=15个
	fmt.Println(h2)
	fmt.Printf("%T\n", h2)
	h3 := []byte(h)
	fmt.Printf("%T\n", h3) //uint8
	h4 := []rune(h)
	fmt.Printf("%T\n", h4) //int32
}

//转义符
/*
\r 回车符
\n 换行符
\t 制表符（水平制表）
\' 单引号
\" 双引号
\\ 反斜杠

*/
//byte:uint8的别称，主要用于区分字节和无符号整型(byte表示一个字节,可以表示英文字符等占用一个字节的字符，多于一个无法正确显示)
//rune:int32的别称，主要用于区分字符和整型（表示字符，用于表示任何一个字符）
/*
byte类型：uint8类型，代表了ASCII字符
rune类型：int32类型，代表了utf-8字符
golang是用utf-8编码的
字节：一字节=8bit
（1）1个ASCII的字符=1个字节：
例如：1个字符'h'=1字节
（2）1个utf-8字符使用1到4个字节为每个字符编码：
utf-8中汉字一般占3个字节(go语言就是用这种格式编译的)
Go语言的字符串的字符的字节使用utf-8编码标识Unicode文本
*/
