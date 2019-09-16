package main

import (
	"fmt"
	"os"
)

var pf = fmt.Printf

type point struct {
	x, y int
}

func main(){
	p := point{1,3}
	pf("%v\n", p)

	pf("%+v\n", p)

	fmt.Printf("%#v\n", p)

	fmt.Printf("%T\n", p)

	fmt.Printf("%t\n", true)

	fmt.Printf("%d\n", 123)

	fmt.Printf("%b\n", 14)

	fmt.Printf("%q\n", "\"string\"")

	fmt.Printf("%x\n", "hex this")

	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")

}