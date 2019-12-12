package main

import (
	"fmt"
	"os"
)

func main() {
	a, b := make(chan int, 3), make(chan int)
	go func() {
		v, ok, s := 0, false, ""

		for {
			select {
			case v, ok = <-a:
				s = "a"
			case v, ok = <-b:
				s = "b"
			}

			if ok {
				fmt.Println(s, v)
			} else {
				os.Exit(0)
			}
		}
	}()

	for index := 0; index < 5; index++ {
		select {
		case a <- index:
		case b <- index:
		}
	}

	close(a)
	select {}
}
