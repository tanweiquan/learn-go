package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//文件操作
func f1() {
	var fileObj *os.File
	var err error
	//打开文件
	fileObj, err = os.Open("./main.go") // 符号./是同一文件夹下目录
	//defer fileObj.Close()//不能在这里写，因为假如err有返回值，则fileObj是nil，nil不能调用Close方法,会导致无法关闭
	if err != nil {
		fmt.Printf("open file faile,err:%v", err)
		return
	}
	//关闭文件
	//defer后面跟着的语句默认是按整个程序文件从上往下每行执行的顺序执行好的，然后由defer语句再放最后执行的
	defer fileObj.Close() //这里要在这里写
}

//读文件方法一
//file.Read()读文件
//格式：func (f *File)Read(b []byte)(n int,err error)
//它接收一个字节切片，返回读取的字节数和可能的具体错误，读到文件末尾是会返回0和io.EOF

func f2() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("ope failed!,err:", err)
		return
	}
	defer file.Close()
	tmp := make([]byte, 128)
	n, err := file.Read(tmp)
	if err == io.EOF {
		fmt.Println("文件读完了")
		return
	}
	if err != nil {
		fmt.Println("read file failed,err:", err)
		return
	}
	fmt.Printf("读取了%d字节数据\n", n)
	fmt.Println(string(tmp[:n]))
}

//for循环读取
func f3() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed!,err:", err)
		return
	}
	defer file.Close()
	var content []byte
	var tmp = make([]byte, 128)
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed,err:", err)
			return
		}
		content = append(content, tmp[:n]...)
	}
	fmt.Println(string(content))
}

//读文件方法二
//bufio读取文件(磁盘-缓存区(buf)-代码)
//bufio是在file的基础上封装了一层API，支持更多的功能
func f4() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file) //创建缓存区读的对象
	for {
		line, err := reader.ReadString('\n') //注意是字符
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed,err:", err)
			return
		}
		fmt.Print(line)
	}
}

//读文件方法三
//ioutil读取整个文件
//io/ioutil包的ReadFile方法能够读取完整的文件，只需将文件名作为参数传入
func f5() {
	content, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Println("read file failed,err:", err)
		return
	}
	fmt.Println(string(content))
}

//文件写入操作
//os.OpenFile()函数能够以指定模式打开文件，从而实现文件写入相关功能
/*格式：
func OpenFile(name string,flag int,perm FileMode)(*file,error){

}
*/
/*其中name:要打开的文件名；flag:打开文件的模式；模式有以下几种
只写：os.O_WRONLY
创建文件：os.O_CREATE
只读：os.O_RDONLY
读写：os.O_RDWR
清空：os.O_TRUNC
追加：os.O_APPEND
perm：文件权限，一个八进制数，r(读)o4,w(写)o2，x(执行)o1。
*/
//文件写入方法一
//Writr和WriteStirng
func f6() {
	file, err := os.OpenFile("xx.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	defer file.Close()
	str := "hello 沙河"
	file.Write([]byte(str))       //写入字节切片数据
	file.WriteString("hello 小王子") //直接写入字符串数据
}

//文件写入方法二
//bufio.NewWriter
func f7() {
	file, err := os.OpenFile("xx.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed,err", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello 沙河\n") //将数据先写入缓存
	}
	writer.Flush() //将缓存中的内容写入文件
}

//文件写入方法三
//ioutil.WriteFile
func f8() {
	str := "hello 沙河"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed,err:", err)
		return
	}
}

func main() {
	f1()
	f2()
	f3()
	f4()
	f5()
	f6()
	f7()
	f8()
}
