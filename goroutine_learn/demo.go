package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(2)
	wg := new(sync.WaitGroup)
	wg.Add(2)

	for index := 0; index < 2; index++ {
		go func(id int) {
			fmt.Println("demoooooooooo")
			defer wg.Done()
			sum(id)
		}(index)
	}

	wg.Wait()
}

func sum(id int) {
	var x int64
	fmt.Println("我想看看MaxUint是啥：", math.MaxUint32)
	fmt.Println(id, x)
	for index := 0; index < math.MaxUint32; index++ {
		x += int64(index)
	}
}
