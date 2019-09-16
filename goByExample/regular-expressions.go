package main

import (
	"fmt"
	"regexp"
)

func main(){
	//match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	//fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")
	//fmt.Println(r.MatchString("peach"))
	//fmt.Println(r.FindAllString("peach punch", -1))
	//fmt.Println(r.FindStringSubmatch("peach punch"))
	//fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	//fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	in := []byte("a peach asdadasd")
	re := []byte("===")
	out := r.ReplaceAll(in, re)
	fmt.Println(string(out))

}
