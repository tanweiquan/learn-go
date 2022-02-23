package main

import (
	"github.com/gin-gonic/gin"
)

/*
 GET用来获取资源
POST用来新建资源
PUT用来更新资源
DELETE用来删除资源。
*/

func main() {
	// 1.新建一个不带中间件的路由
	r := gin.New()

	// 2.绑定路由规则，执行函数

	// gin.Context,封装了request和response
	// GET：请求方式；/hello1：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.GET("/hello1", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		// gin.H相当于在使用gin包里的type H map[string]interface{}
		c.JSON(200, gin.H{"message": "Hello GET method!"}) //这里的200是响应状态码
	})

	// POST：请求方式；/hello2：请求的路径
	// 当客户端以POST方法请求/hello2路径时，会执行后面的匿名函数
	r.POST("/hello2", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello POST method!"})
	})

	// PUT：请求方式；/hello3：请求的路径
	// 当客户端以PUT方法请求/hello3路径时，会执行后面的匿名函数
	r.PUT("/hello3", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello PUT method!"})
	})

	// DELETE：请求方式；/hello4：请求的路径
	// 当客户端以DELETE方法请求/hello4路径时，会执行后面的匿名函数
	r.DELETE("/hello4", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello DELETE method!"})
	})

	// 3.启动HTTP服务，默认在0.0.0.0:8080启动服务
	// _ = r.Run(":8080")
	_ = r.Run()
}
