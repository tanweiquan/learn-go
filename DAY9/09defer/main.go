package main

import "fmt"

/*
defer语句，defer语句会将其后面跟随的语句进行延迟处理，
在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行，
也就是说，先被defer的语句最后被执行，最后被defer的语句，最后被执行。
*/
//所有语句都是按正常顺序从上往下跑的，只是前面带有defer的语句，去defer后延缓执行而已(defer+语句一起构成了新语句)
func f1(x string, a, b int) int {
	tmp := a + b
	fmt.Println(x, a, b, tmp)
	return tmp
}

// defer 是后进先出，panic 需要等defer 结束后才会向上传递。
// 即如果发生painc，崩溃前会执行完全部defer登记的代码
func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
}
func main() {
	defer_call() //打印后->打印中->打印前->触发异常
	//fmt.Println("start")
	//下面打印结果是4，5，3，2，1。也就是说defer延迟处理语句时，defer跟随的语句，在其他语句执行完后再被执行
	//defer fmt.Println(1) //最后被执行
	//defer fmt.Println(2)
	//defer fmt.Println(3)
	//fmt.Println(4)
	//defer fmt.Println(5) //最先被执行（也就是说先defer语句最后被执行，最后defer语句行最先被执行，但都晚执行）
	a := 1
	b := 2
	defer f1("1", a, f1("10", a, b))
	a = 0
	defer f1("2", a, f1("20", a, b))
	b = 1
	//下面解释上面的执行步骤
	//1.a := 1
	//2.b := 2
	//3.defer f1("1",1,f1("10",1,2))
	//4.f1("10",1,2)//10 1 2 3
	//5.defer f1("1",1,3)
	//6.a=0
	//7.defer f1("2",0,f1("20",0,2))
	//8.f1("20",0,2)//20 0 2 2
	//9.defer f1("2",0,2)
	//10.b=1
	//11.f1("2",0,2)//按f1的意思要打印出x,a,b,tmp的值，即打印结果为2 0 2 2
	//12.f1("1",1,3)//按f1的意思要打印出x,a,b,tmp的值，即打印结果为1 1 3 4

}
