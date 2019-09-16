package main

import (
	"fmt"
)

func main() {
	// var a interface{} = 2.0123123
	// fmt.Print(a, "\n\n")
	// goHome(22)
	// mapDemo()
	// mapRange()
	forPrint()
}

func forPrint() {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		fmt.Println(len(s), cap(s))
	}
}

func mapRange() {
	hash := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
	}
	for k, v := range hash {
		fmt.Printf("%v, %v \n", k, v)
	}
}

func mapDemo() {
	arr := []int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%v, %v \n", i, arr[i])
		arr = append(arr, arr[i])
	}
	fmt.Println(arr)
}

func goHome(a int) bool {
	return a >= 21
}

func Tag(tag int) {
	switch tag {
	case 1:
		fmt.Print("andriod")
	case 2:
		fmt.Print("go")
	case 3:
		fmt.Print("java")
	default:
		fmt.Print("c")
	}
}
