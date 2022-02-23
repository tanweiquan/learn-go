package main

import (
	"fmt"
	"time"
)

//时间
//time.Time（前面是time包，后面是Time类型）
//时间戳：时间戳是自1970年1月1日(08:00:00GMT)至当前时间的总毫秒数。别称为Unix时间戳(UnixTimestamp)

func main() {
	now := time.Now()
	fmt.Println(now)          //2021-12-25 17:06:04.6487271 +0800 CST m=+0.005260401
	fmt.Println(now.Date())   //日期
	fmt.Println(now.Year())   //年
	fmt.Println(now.Month())  //月
	fmt.Println(now.Day())    //日
	fmt.Println(now.Hour())   //时
	fmt.Println(now.Minute()) //分
	fmt.Println(now.Second()) //秒
	//时间戳
	timeStamp1 := now.Unix()     //时间戳
	timeStamp2 := now.UnixNano() //纳秒时间戳
	fmt.Println(timeStamp1)
	fmt.Println(timeStamp2)
	//time.Unix()还可以将时间戳转化为时间格式
	ret := time.Unix(timeStamp1, 0)
	fmt.Println(ret)
	//时间间隔
	//Duration类型代表两个时间点之间经过的时间,以纳秒为单位
	/*time包定义的时间间隔常量如下：
	const (
		Nanosecond  Duration = 1
		Microsecond          = 1000 * Nanosecond
		Millisecond          = 1000 * Microsecond
		Second               = 1000 * Millisecond
		Minute               = 60 * Second
		Hour                 = 60 * Minute
	)
	例如：time.Duration表示1纳秒，time.Second表示1秒
	*/
	//时间+时间间隔要求：
	//go语言的时间对象有提供Add方法
	later := now.Add(time.Hour)         //later被赋值为1小时后的时间
	fmt.Println(later)                  //打印结果为1小时后的时间
	_24later := now.Add(24 * time.Hour) //_24later被赋值为24小时后的时间
	fmt.Println(_24later)
	//定时器
	// timer := time.Tick(time.Second) //每隔1秒将时间返回给timer
	// for t := range timer {
	// 	fmt.Println(t) //1秒执行一次
	// }
	//时间格式化：把语言中时间对象转换成字符串类型的时间
	//时间类型有一个自带的方法Format进行格式化，时间模板为2006年1月2日15点04分，记忆口诀为2006-1-2-3-4-5
	//将时间转换为2006-01-02-03-04的格式
	fmt.Println(now.Format("2006-01-02-15-04")) //打印结果为2021-12-25-17-49
	//转换为2006/01/02/03/04/05.000的格式
	fmt.Println(now.Format("2006/01/02/15/04/05.000")) //打印结果为2021/12/25/18/00/46.249
	//Sleep
	fmt.Println("开始Sleep了")
	time.Sleep(5 * time.Second) //停五秒钟
	fmt.Println("五秒过去了")
}
