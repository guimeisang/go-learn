/*
golang  第 1 次作业：
1、完成 go 环境搭建，设置 GOPATH，选一个你喜欢的 IDE（推荐 jetbrains 的 Goland 或者 VSCode）
2、使用 fmt 包的 Println 打印 "Hello, golang"
> 弄清楚了 %v, %T, %+v, %#v, %d, %[n]d, %*d

3、写出 GOPATH 的作用
> 为什么需要设置每一个项目的GOPATH, 因为GOPATH是影响import的核心

4、使用 go build、fmt、run 命令
> go build 是编译
> go fmt 是格式化代码
> go run 是运行go文件

5、说说你搜到的 plan9 汇编的资料



*/

// 完成 go 环境搭建，设置 GOPATH，选一个你喜欢的 IDE（推荐 jetbrains 的 Goland 或者 VSCode）
// 整个比较难的还是配置GOPATH, 后面知道可以用export配置就基本上没啥了

// 使用 fmt 包的 Println 打印 "Hello, golang"

package main

import "testing"

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Error("ashdjad")
			main()
		})
	}
}
