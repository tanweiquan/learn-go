package split_string

//我们在split_string包中定义了一个Split函数，具体实现如下：

// split_string/split.go
import (
	"strings"
)

// Split 把字符串s按照给定的分隔符sep进行分割返回字符串切片
func Split(s, sep string) (result []string) {
	i := strings.Index(s, sep)

	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):]
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}

func Hello(n int) int {
	if n < 2 {
		return n
	}
	return Hello(n-1) + Hello(n-2)
}

//在目录文件下输入go test -cover可以检测函数测试覆盖率(即用例测试函数的语句含整个函数的占比)
//输出的coverage: xx.x% of statements就是函数测试覆盖率
