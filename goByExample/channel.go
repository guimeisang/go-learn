package main

import "fmt"

func main() {
	msg := make(chan string)

	go func() {
		msg <- "ping"
	}()

	_msg := <-msg
	fmt.Println(_msg)
}
