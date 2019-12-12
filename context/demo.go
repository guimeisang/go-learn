package main

import (
	"context"
	"fmt"
	"time"
)

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				//执行defer cancel操作后，就会执行到该select入库
				fmt.Println("i exited")
				return // returning not to leak the goroutine
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

func test() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	intChan := gen(ctx)

	for n := range intChan {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

func main() {
	test()
	time.Sleep(time.Second * 5)
}
