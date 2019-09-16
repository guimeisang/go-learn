package main

import "time"
import "fmt"

func main() {
	c1 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "one"
	}()

	select {
	case msg1 := <-c1:
		fmt.Println(msg1)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout")
	}

	c2 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	select {
	case msg2 := <-c2:
		fmt.Println(msg2)
	case <-time.After(time.Second * 3):
		fmt.Println("two timeout")
	}
}
