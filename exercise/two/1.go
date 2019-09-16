/**
Go 训练营第二次作业

总体来说第二次作业，还是比较基础的

1、使用 Go 语言 var 关键字声明 Go 中所有的数据类型，并使用 fmt 打印类型名

2、「类型推导」变量声明并初始化有哪几种方式？写出示例代码

3、如何用一行代码交换整数 i 和 j 的值

4、Go 中整型有哪几种类型，这些类型 Java 中是否支持？
int 
uint
uintptr
int32 
int64
int8
uint8
uint32
uint64

var a int32
value1 := 64
value2 = int64(a)


5、Go 的浮点数有哪几种类型，这些类型 Java 中是否支持？
float32和float64

1. %g 打印的格式
2. 默认是float64: value := 1.01
3. 强制转换类型：float32(value)
4. 比较两个浮点数大小方式：由于浮点数本来就不是很准确的，所以采用如下的方式进行比较：

// 由用户去决定精度是怎么样的
func isEqual(f1, f2, p float64) bool {
	return math.Fdim(f1 - f2) < p
}


参考文档:https://studygolang.com/articles/6566

*/

