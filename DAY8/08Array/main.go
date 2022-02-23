package main

import "fmt"

//数组Array
//存放元素的容器
//必须指定存放元素的类型和容量
//数组的长度是数组类型的一部分
func main() {
	var a [3]bool //[flase flase flase]
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	//数组初始化
	var b [3]int                        //初始化b为int类型的零值
	var c = [3]int{1, 2}                //使用指定的元素完成初始化,其中类型前的是长度
	var e = [3]string{"广州", "深圳", "成都"} //使用指定的元素完成初始化
	fmt.Println(b, c, e)
	//数组如果不初始化，元素会默认是零值（布尔值是flase，整型和浮点型是0，字符串是""）
	//初始化方式
	a1 := [5]int{0, 2, 3, 5, 1}
	a2 := [...]int{1, 3, 5, 2, 1, 5} //根据数组自行判断长度
	a3 := [6]int{0: 1, 5: 6}         //根据索引来初始化,改变了不同位置的值
	fmt.Println(a1, a2, a3)
	//数组的遍历
	h := [...]string{"广州", "长沙", "成都", "深圳"}
	//1.根据索引遍历
	for i := 0; i < len(h); i++ {
		fmt.Println(h[i]) //h[i]是代表数组h里内位置为i的元素，其中位置是从0开始算起的
	}
	//2.for range遍历
	for i, v := range h {
		fmt.Println(i, v)
	}
	//多维数组
	//例如，[[1,2],[2,4],[3,4]]

	a4 := [3][2]int{
		{1, 2},
		{2, 4},
		{3, 4},
	}
	fmt.Println(a4)
	//多维数组的遍历
	for _, v1 := range a4 {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
	//数组是值类型
	b1 := [3]int{1, 2, 3}
	b2 := b1 //相当于复制过程，把文档从A文件夹复制到B文件夹
	b2[0] = 100
	fmt.Println(b2) //b2:[100,2,3]
	fmt.Println(b1) //b1:[1,2,3]
	c1 := [3]int{4, 5, 3}
	fmt.Printf("c1修改前=%d", c1) //[4 5 3],不影响修改前的数组里的元素
	c1[1] = 12
	fmt.Printf("c1修改后=%d", c1) //[4 12 3]，修改后就是把对应元素的内存地址更改了
}
