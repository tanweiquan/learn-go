package main

//下载第三方redis
//go get github.com/go-redis/redis
//假如下载不成功,则要先配置代理GOPROXY=???，再执行go get
/* go get 参数
-d 只下载不安装
-f 只有在你包含了 -u 参数的时候才有效，不让 -u 去验证 import 中的每一个都已经获取了，这对于本地 fork 的包特别有用
-fix 在获取源码之后先运行 fix，然后再去做其他的事情
-t 同时也下载需要为运行测试所需要的包
-u 强制使用网络去更新包和它的依赖包，即更新本地已有的包及其依赖
-v 显示执行的命令 */
import (
	"fmt"

	"github.com/go-redis/redis"
)

var redisDb *redis.Client //声明是一个client,有类似mysql的连接池

func initRedis() (err error) {
	redisDb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", //要连接的服务端的ip及端口
		Password: "",               //密码
		DB:       0,                //redis要使用多少种数据库
		PoolSize: 100,              // 连接池大小
	})
	_, err = redisDb.Ping().Result()
	if err != nil {
		fmt.Printf("ping failed,err:%v\n", err)
		return
	}
	return
}
func main() {

	//启动客户端(运行前要启动服务端先)
	err := initRedis()
	if err != nil {
		fmt.Printf("connect failed,err:%v\n", err)
		return
	}
	fmt.Println("连接成功！")

	//现在可以调用redis包里的api了
	//zset
	zsetKey := "language_rank"
	languages := []redis.Z{
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 97.0, Member: "JavaScript"},
		{Score: 99.0, Member: "C/C++"},
	}
	// ZADD
	num, err := redisDb.ZAdd(zsetKey, languages...).Result()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Printf("zadd %d succ.\n", num)
	// 把Golang的分数加10
	newScore, err := redisDb.ZIncrBy(zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)

	// 取分数最高的3个
	ret, err := redisDb.ZRevRangeWithScores(zsetKey, 0, 2).Result()
	if err != nil {
		fmt.Printf("zrevrange failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
	// 取95~100分的
	op := redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = redisDb.ZRangeByScoreWithScores(zsetKey, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}
