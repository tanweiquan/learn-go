package main

import "fmt"

//递归：函数自己调用自己
//递归适合处理那种问题相同，相同问题规模越来越小的场景
//递归一定要有一个明确的退出条件
/*形式如下：
func f1() {
	f1()
}
*/
//计算n的阶乘
func f(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * f(n-1) //自己调用自己
}

//面试题：n个台阶，一次可以走1步，也可以走2步，有多少种走法
//递归实现如下：
func taijie(n uint64) uint64 {
	//如果只有一个台阶就一种走法：
	if n == 1 {
		return 1
	}
	//如果只有两个台阶就两种走法：
	if n == 2 {
		return 2
	}
	//多个台阶的走法：(剩余一个台阶+剩余两个台阶)
	return taijie(n-1) + taijie(n-2)
}
func main() {
	//计算5的阶乘 5!=5*4*3*2*1
	ret := f(5)
	fmt.Println(ret)
	//计算4个台阶的走法
	ret1 := taijie(4)
	fmt.Println(ret1)
}
