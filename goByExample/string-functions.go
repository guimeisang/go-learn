package main

import "fmt"
import str "strings"  // 这个就是在引入包的时候还可以增加别名

var p = fmt.Println //小技能，可以少打字

func main(){
	p("Contains:  ", str.Contains("test", "es"))
	p("Count:     ", str.Count("test", "t"))
	p("HasPrefix: ", str.HasPrefix("test", "te"))
	p("HasSuffix: ", str.HasSuffix("test", "st"))
	p("Index:     ", str.Index("test", "e"))
	p("Join:      ", str.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", str.Repeat("a", 5))
	p("Replace:   ", str.Replace("foo", "o", "0", -1))
	p("Replace:   ", str.Replace("foo", "o", "0", 1))
	p("Split:     ", str.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", str.ToLower("TEST"))
	p("ToUpper:   ", str.ToUpper("test"))
}
