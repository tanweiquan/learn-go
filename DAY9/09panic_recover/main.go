package main

//panic以及recover
//panic导致程序崩溃
//1.recover()必须要搭配defer使用
//2.defer一定要在可能引发panic的语句之前定义(也就是defer语句一定要在可能发生错误崩溃的语句前使用)
import "fmt"

func f1() {
	fmt.Println("a")
}
func f2() {
	//假如刚刚打开数据库连接
	defer func() {
		fmt.Println("释放数据库连接")
		err := recover() //这里用recover将错误传到变量err中
		fmt.Println(err)
	}()
	panic("程序出现严重错误！！！") //程序崩溃退出
	fmt.Println("b")
}
func f3() {
	fmt.Println("c")
}
func main() {
	f1() //a
	f2() //释放数据库连接 程序出现严重错误！！！
	f3() //c
}
