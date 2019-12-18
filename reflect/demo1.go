package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	t := reflect.TypeOf(x)
	fmt.Println("type: ", t)
	v := reflect.ValueOf(x)
	fmt.Println("value: ", v)
	fmt.Println("value: ", v.Type())
	fmt.Println("value: ", v.Kind())
	fmt.Println("value: ", v.Float())
	fmt.Println("value: ", v.Interface())
	y := v.Interface().(float64)
	fmt.Println(y)
}
