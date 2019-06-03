package main

import "fmt"

func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "int")
		case string:
			fmt.Println(arg, "int")
		case int64:
			fmt.Println(arg, "int")
		default:
			fmt.Println(arg, "unKnow")
		}
	}
}
