package main

import (
	"flag"
	"time"
)

//flag获取命令行参数
//go run main.go -name=沙河 -age=10 -boolValue=true -clientTime=10h
//13flag.exe -name=沙河 -age=10 -boolValue=true -clientTime=10h
/*
如：13flag.exe -name=沙河 -age=10 -boolValue=true -clientTime=10h a b c
flag.Args()  //返回命令行参数后的其他参数，以[]string类型 //[a b c]
flag.NArg()  //返回命令行参数后的其他参数个数 //3
flag.NFlag() //返回使用的命令行参数个数 //4
*/
func main() {
	/*
		//得到的是一个指针（内存地址）
		name := flag.String("name", "默认名", "请输入名字")
		age := flag.Int("age", 0, "请输入年龄")
		boolValue := flag.Bool("boolValue", false, "请输入判断信息")
		clientTime := flag.Duration("clientTime", time.Second, "请输入时间")
		//解析flag
		flag.Parse()
		fmt.Println(*name)
		fmt.Println(*age)
		fmt.Println(*boolValue)
		fmt.Println(*clientTime)
		fmt.Println(flag.Args())
		fmt.Println(flag.NArg())
		fmt.Println(flag.NFlag())

	*/
	var name string
	var age int
	var boolValue bool
	var clientTime time.Duration
	flag.StringVar(&name, "name", "默认名", "请输入名字")
	flag.IntVar(&age, "age", 0, "请输入年龄")
	flag.BoolVar(&boolValue, "boolValue", false, "请输入判断信息")
	flag.DurationVar(&clientTime, "clientTime", time.Second, "请输入时间")
	//另外还有pprof性能分析工具
}
