module learn-go

go 1.16

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-sql-driver/mysql v1.6.0
	github.com/jmoiron/sqlx v1.3.4
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.18.1 // indirect
	github.com/streadway/amqp v1.0.0
)

// go mod init       # 初始化modules
// go mod tidy       # 检查，删除错误或者不使用的modules，下载没download的package
// go mod download   # 下载modules到本地cache
// go mod vendor     # 生成vendor目录,将依赖转移至本地的vendor文件
// go mod edit       # 编辑go.mod文件，选项有-json、-require和-exclude，可以使用帮助go help mod edit
// go mod graph      # 打印依赖图
// go mod verify     # 验证依赖是否正确
// go mod why        # 查找依赖
// go mod graph      # 以文本模式打印模块需求图
// go test           # 执行一下，自动导包
