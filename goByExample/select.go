package main

import "time"
import "fmt"

func A(a chan string) {
	time.Sleep(time.Second * 1)
	a <- "one"
}

func B(b chan string) {
	time.Sleep(time.Second * 2)
	b <- "two"
}

func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	go A(c1)
	go B(c2)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
