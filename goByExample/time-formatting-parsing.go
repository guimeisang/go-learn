package main

import (
	"fmt"
	"time"
)

func main(){
	p := fmt.Println
	t := time.Now()

	p(t.Format(time.RFC3339))

	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)
	p(e)
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))

	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)

}
