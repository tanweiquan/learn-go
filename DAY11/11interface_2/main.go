package main

import "fmt"

//接口示例类型2
//不管是什么牌子的车，都能跑

//定义一个car接口类型
//不管什么结构体，只要有run方法都是car类型
type car interface {
	run()
}

//造不同牌子的车
type falali struct {
	brand string
}

type baoshijie struct {
	brand string
}

//造不同车的方法
func (f falali) run() {
	fmt.Printf("%s速度70迈~", f.brand)
}

func (b baoshijie) run() {
	fmt.Printf("%s速度75迈~", b.brand)
}

//drive函数接收一个car类型的变量
func drive(c car) {
	c.run()
}

func main() {
	var f1 = falali{
		brand: "法拉利",
	}
	var b1 = baoshijie{
		brand: "保时捷",
	}
	drive(f1)
	drive(b1)

}
