package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string
	Age   int
	Score float32
}

func (s Student) Print() {
	fmt.Println(s)
}

func (s Student) Set(name string, age int, score float32) {
	s.Name = name
	s.Age = age
	s.Score = score
}

func TestStruct(a interface{}) {
	v := reflect.ValueOf(a)
	kd := v.Kind()

	fmt.Println(v, kd)

	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	// 获取字段数量
	fields := v.NumField()
	fmt.Printf("struct has %d field\n", fields)
	for index := 0; index < fields; index++ {
		kd := v.Field(index).Kind()
		fmt.Printf("%d %v\n", index, kd)
	}

	// 获取方法数量
	methods := v.NumMethod()
	fmt.Printf("struct has %d methods\n", methods)

	// 放射调用的Print方法
	var params []reflect.Value
	v.Method(0).Call(params)
}

func main() {
	var a Student = Student{
		Name:  "stu01",
		Age:   18,
		Score: 98,
	}

	TestStruct(a)
}
