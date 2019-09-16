package main

import (
	"fmt"
	"sort"
)

func main(){
	var strs = []string{"b", "a", "c"}
	sort.Strings(strs)
	fmt.Println(strs)

	ints := []int{7, 3, 4}
	sort.Ints(ints)
	fmt.Println(ints)

	isSort := sort.IntsAreSorted(ints)
	fmt.Println(isSort)
}
