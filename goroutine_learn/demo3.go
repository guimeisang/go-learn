package main

import (
	"fmt"
	"time"
)

func test() {
	fmt.Println("test")
}

func calc() {
	fmt.Println("calc")
}

func main() {
	go test()
	for index := 0; index < 2; index++ {
		go calc()
	}
	time.Sleep(10 * time.Second)
}
