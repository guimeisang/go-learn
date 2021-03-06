package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string `json: "stu_name"`
	Age   int
	Score float32
}

func TestStruct(a interface{}) {
	typ := reflect.TypeOf(a)

	tag := typ.Elem().Field(0).Tag.Get("json")
	fmt.Printf("Tag:%s\n", tag)
}

func main() {
	var a Student = Student{
		Name:  "stu01",
		Age:   18,
		Score: 92.8,
	}
	TestStruct(&a)
}
