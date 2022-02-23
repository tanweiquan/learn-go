package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// net/http client

func main() {
	/* resp, err := http.Get("http://127.0.0.1:9090/hello/?name=沙河&age=18")
	if err != nil {
		fmt.Printf("Get url failed,err:%v\n", err)
		return
	} */

	//客户端解析服务端网址
	data := url.Values{} //url Values
	urlObj, _ := url.Parse("http://127.0.0.1:9090/hello/")

	data.Set("name", "沙河")
	data.Set("age", "18")
	queryStr := data.Encode() //得到encode之后的URL
	fmt.Println(queryStr)
	urlObj.RawQuery = queryStr
	req, err := http.NewRequest("GET", urlObj.String(), nil)
	if err != nil {
		fmt.Printf("Get url failed,err:%v\n", err)
		return
	}

	//发送http请求。req是具体的请求对象；resp是返回的http响应
	resp, err := http.DefaultClient.Do(req) //DefaultClient是默认客户端
	if err != nil {
		fmt.Printf("Get url failed,err:%v\n", err)
		return
	}

	//从resp中把服务端返回的数据读出来
	//var data []byte
	//resp.Body.Read()
	//resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body) //我在客户端读出服务端返回的响应的body
	if err != nil {
		fmt.Printf("read resp.Body failed,err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}
