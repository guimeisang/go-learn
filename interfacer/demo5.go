package main

import "fmt"

type Sayer interface {
	Say(message string)
	SayHi()
}

type Animal struct {
	Name string
}

func (a *Animal) Say(message string) {
	fmt.Println("Animal[%v] say: %v\n", a.Name, message)
}

func (a *Animal) SayHi() {
	a.Say("Hi")
}

type Dog struct {
	Animal
}

func (d *Dog) Say(message string) {
	fmt.Println("Dog[%v] say: %v\n", d.Name, message)
}

type Cat struct {
	Animal
}

func (c *Cat) Say(message string) {
	fmt.Printf("Cat[%v] say: %v\n", c.Name, message)
}

func (c *Cat) SayHi() {
	c.Say("hi")
}

func main() {
	var sayer sayer

	sayer = &Cat{Animal{Name: "Jerry"}}
	sayer.Say("hello world")
	sayer.SayHi()
}
