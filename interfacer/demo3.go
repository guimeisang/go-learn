package main

import "fmt"

type User struct {
	id   int
	name string
}

func (self *User) String() string {
	return fmt.Sprintf("%d, %s", self.id, self.name)
}

func main() {
	var o interface{} = &User{1, "Tom"}
	switch v := o.(type) {
	case nil:
		fmt.Println("nil")
	case fmt.Stringer:
		fmt.Println(v)
	case func() string:
		fmt.Println(v())
	case *User:
		fmt.Println("%d, %s\n", v.id, v.name)
	default:
		fmt.Println("unknow")
	}
}
