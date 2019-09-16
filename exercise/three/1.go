package main

/**
1、如何 build 不同平台的可执行包？比如把 main.go 在 Mac 下构建 windows、linux 的可执行包

GOOS：目标可执行程序运行操作系统，支持 darwin，freebsd，linux，windows
GOARCH：目标可执行程序操作系统构架，包括 386，amd64，arm

就可以了


2、Go 的单元测试怎么写
- 文件名规则：*_test.go 这样go的语言工具才能识别
- 函数名规则：Test*() 这样可以暴露出去
- 函数中必须接受一个指向 *testing.T 类型的指针，并且不返回任何东西
- 测试不ok，用t.Fatal, 正确的用t.Log
- Mock 函数：适当的时候需要使用httptest.NewServer这种方式提供
- 覆盖率：go test -v -coverprofile=a.out || go tool cover -html=c.out -o=a.html 用浏览器打开就可以看到代码覆盖的情况，根据覆盖的情况增加单元测试；

参考文档：https://www.flysnow.org/2017/05/16/go-in-action-go-unit-test.html


3、如何遍历数组
GOOS=linux GOARCH=amd64 go tool compile -S main.go
go build -gcflags -S  main.go

上面是可以看到main.go 的汇编，可以看到range和for都汇编成同一个东西：





4、写一个数组切片的例子，说说切片len 和 cap 的区别





5、有这样的数组：
arr3 := [...]int{1, 2, 3, 4, 5}
如何把它切成前 3 个数一个数组，后 2 个数一个数组



6、 下面的代码输出什么？
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		println(len(s), cap(s))
	}

输出 1-10没有增加，但是另外一个却一直在增加很奇怪



7、如何声明 map，map 如何遍历





8、如何声明一个 key 为 string，value 可以为任意类型的 map？






*/
