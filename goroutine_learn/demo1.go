package main

import "fmt"

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	ch := make(chan int)
	go sum(a[:len(a)/2], ch)
	go sum(a[len(a)/2:], ch)

	fmt.Println("ch 中没有数据我将阻塞起等数据")
	x, y := <-ch, <-ch
	fmt.Println(x, y, x+y)
}

func sum(a []int, ch chan int) {
	var sum = 0
	for _, v := range a {
		sum += v
	}

	ch <- sum
}
