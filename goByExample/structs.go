package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 20})

	fmt.Println(person{name: "fred"})

	fmt.Println(&person{name: "Ann", age: 20})

	s := person{name: "Ann", age: 20}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 15
	fmt.Println(sp.age)
}
