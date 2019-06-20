package main

import "fmt"

func f(a string) {
	for i := 0; i < 3; i++ {
		fmt.Println(a)
	}
}

func main() {
	f("direct")

	go f("gorountine")

	go func(msg string) {
		fmt.Println(msg)
	}("asda")
	// var input string
	// fmt.Scanln(&input)
	fmt.Println("done")
}
