package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		defer wg.Done()
		defer fmt.Println("A defer")

		func() {
			defer fmt.Println("B defer")
			runtime.Goexit()
			fmt.Println("B")
		}()

		fmt.Println("哈哈哈")
	}()

	wg.Wait()
}
