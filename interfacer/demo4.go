package main

import "fmt"

type Data struct {
	id   int
	name string
}

func (self *Data) String() string {
	return fmt.Sprintf("%d, %s", self.id, self.name)
}

func main() {
	var _ fmt.Stringer = (*Data)(nil)
}
