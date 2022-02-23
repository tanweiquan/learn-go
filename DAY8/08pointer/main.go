package main

import "fmt"

//vscode 不支持go module
//pointer
//go语言中，指针是一种类型，指向变量所在的内存单元(不是内存地址)
//注意：变量名前加上星号字符，比如*age，指向变量age所在的内存单元
//在go语言中，目前的引用类型有：切片、map、chan、func
//简单来说，值为内存地址的变量是指针
//假如有b:=&p,则b和&p都是指针。假如有*a,则a是指针。假如有声明a类型为*p，则a为指针类型，*p是一种类型。

func main() {
	//1、&是取地址符号，即取得某个变量的地址，如：&a
	//2、*是指针运算符，可以表示一个变量是指针类型，也可以表示一个指针变量所指向的内存储存单元，也就是这个地址所储存的值
	//假如某个变量的值是内存地址，则该变量是指针类型。如：b:=&a，则b是指针类型；
	//假如声明了变量是指针类型，则该变量储存的值是一个内存地址。如：var a *int,b=10,a=&b

	a := 18
	fmt.Println(a)        //18;打印a的值
	fmt.Println(&a)       //0xc000014098;打印a的内存地址
	b := &a               //因为b的值是a的内存地址，因此b是值为a的内存地址的指针
	fmt.Println(b)        //0xc000014098;打印b的值，大小为a的内存地址
	fmt.Println(*b)       //18;打印内存地址为&a的值*b，即打印a的值
	fmt.Printf("%T\n", b) //*int，指针类型
	//a的类型=int &a的类型=*int b的类型=*int *b的类型=int
	fmt.Printf("a的类型=%T &a的类型=%T b的类型=%T *b的类型=%T\n", a, &a, b, *b)
	//*b = 100//这样做相当于把之前的a的值和内存地址都改变了
	//fmt.Println(*b, b)//打印值和内存都变了
	//1、make和new都是用来申请内存的
	//2、new很少用，一般用来给基本数据类型申请内存，string\int，返回的是对应类型的指针
	//3、make是用来给slice、map、channel申请内存的，make函数返回的是对应的这三个类型本身
	//4new函数申请一个内存地址(make也是分配内存地址的，但make只用于slice、map、channel，不用返回指针)
	var s *int      //nil pointer(空指针)
	fmt.Println(s)  //<nil>
	fmt.Println(&s) //0xc000006030
	var c = new(int)
	fmt.Println(c)  //0xc0000140e8
	fmt.Println(*c) //0
	*c = 100        //不影响上面c的内存地址，只改变c的值，相当于在c的内存上多申请一个位置
	fmt.Println(c)  //0xc0000140e8，返回指针
	fmt.Println(*c) //100
}
