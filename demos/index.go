package main

import "fmt"

func modify() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{3, 3, 3}
	copy(slice1, slice2)
	fmt.Println("In modify(), array values:", slice1)
	copy(slice2, slice1)
	fmt.Println("In modify(), array values:", slice2)
}

func main() {
	modify()
}
