package split_string

import (
	"reflect"
	"testing"
)

//单元测试与基准测试
//.................................................
//单元测试

//测试函数覆盖率要达到100%
//测试代码覆盖率要达到60%

//写完测试用例后，在终端中输入go test -v 可以测试多个用例，看到多个用例的测试结果
//测试用例
/* func TestSplit(t *testing.T) {
	got := Split("abcdbls", "b")      // 程序输出的结果
	want := []string{"a", "cd", "ls"} // 期望的结果
	//判断got是否等于want
	if !reflect.DeepEqual(got, want) { // 因为slice不能比较直接，借助反射包中的方法比较
		//测试用例失败了
		t.Errorf("want:%v but got:%v", want, got) // 测试失败输出错误提示
	}
}
func Test2Split(t *testing.T) {
	got := Split("a:b:c:d", ":")
	want := []string{"a", "b", "c", "d"}
	if !reflect.DeepEqual(got, want) {
		//测试用例失败了
		t.Errorf("want:%v but got:%v", want, got)
	}
} */

//测试组
//使用slice切片做测试组(比起以上多个测试用例分开写的优点是可以快速追加测试用例)
func TestSplit(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}
	//声明测试组使用slice类型
	testGroup := []testCase{
		{"abcdbls", "b", []string{"a", "cd", "ls"}},
		{"a:b:c:d", ":", []string{"a", "b", "c", "d"}},
		{"你好，沙河", "，", []string{"你好", "沙河"}},
	}
	for _, ts := range testGroup {
		got := Split(ts.str, ts.sep)
		if !reflect.DeepEqual(got, ts.want) {
			//测试用例失败了
			t.Fatalf("want:%v but got:%v", ts.want, got)
		}
	}
}

//使用map做测试组(比起以上用切片做测试组的优点是可以快速看出哪个测试用例输出有问题)
func Test2Split(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}
	//声明测试组使用map类型
	testGroup := map[string]testCase{
		"case1": {"abcdbls", "b", []string{"a", "cd", "ls"}},
		"case2": {"a:b:c:d", ":", []string{"a", "b", "c", "d"}},
		"case3": {"你好，沙河", "，", []string{"你好", "沙河"}},
	}
	for name, ts := range testGroup {
		got := Split(ts.str, ts.sep)
		if !reflect.DeepEqual(got, ts.want) {
			//测试用例失败了
			t.Fatalf("name:%s want:%v but got:%v", name, ts.want, got)
		}
	}
}

//子测试：使用map类型的测试组的情况下，单独测试某个用例
//在终端中输入go test -run=测试函数+/+用例名
func Test3Split(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}
	//声明测试组使用map类型
	testGroup := map[string]testCase{
		"case1": {"abcdbls", "b", []string{"a", "cd", "ls"}},
		"case2": {"a:b:c:d", ":", []string{"a", "b", "c", "d"}},
		"case3": {"你好，沙河", "，", []string{"你好", "沙河"}},
	}
	for name, ts := range testGroup {
		//测试其中的一个case
		// t.Run(name,func(t func(t *testing.T)))
		t.Run(name, func(t *testing.T) {
			got := Split(ts.str, ts.sep)
			if !reflect.DeepEqual(got, ts.want) {
				//测试用例失败了
				t.Fatalf("want:%v but got:%v", ts.want, got)
			}
		})
	}
}

//在目录文件下输入go test -cover可以检测函数测试覆盖率(即用例测试函数的语句含整个函数的占比)
//输出的coverage: xx.x% of statements就是函数测试覆盖率
//将覆盖率输出到一个文件，输入go test -coverprofile=xx.out  //将覆盖率输出到xx.out这个文件里
//输入go tool cover -html=xx.out //使用html分析测试覆盖率

//.........................................................................................................
//Benchmark函数名 基准测试
//输入go test -bench=函数名 //执行测试Benchmark
//如输入go test -bench=Split,循环测试函数Split一秒种
//假如要强制执行不同时间，可以输入go test -bench=Split -benchtime=20s  //20秒内循环执行函数Split
//指定使用cpu的数量：go test -bench=函数名 -cpu 数目
//指定使用cpu的数量：go test -bench=Split -cpu 1
//输入go test -bench=函数名 -mem //查看循环执行函数1秒，内存的申请情况，从而做性能优化（如让切片做make初始化，减少内存申请次数）
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a,b,c", ",") //传值后测试执行这个函数，某一段时间内能跑多少次跑多少次
	}
}

//测试传入第n个元素后，循环执行1秒的情况
func benchmarkHello(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Hello(n)
	}
}

//运行时，终端输入go test -bench=测试的函数名
//如终端输入go test -bench=Hello2
//运行所有的要输入元素的函数时，输入go test -bench=.
func BenchmarkHello1(b *testing.B) { benchmarkHello(b, 1) } //传入第1个元素循环执行1秒的情况
func BenchmarkHello2(b *testing.B) { benchmarkHello(b, 2) } //传入第2个元素循环执行1秒的情况
func BenchmarkHello3(b *testing.B) { benchmarkHello(b, 3) } //传入第3个元素循环执行1秒的情况
