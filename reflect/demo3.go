package main

import (
	"fmt"
	"reflect"
)

func main() {

	var b int = 1
	b = 200
	testInt(&b)
	fmt.Println(b)
}

func testInt(b interface{}) {
	val := reflect.ValueOf(b)
	val.Elem().SetInt(100)
	fmt.Println(val.Elem())
}
