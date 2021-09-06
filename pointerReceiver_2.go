package main

import (
	"fmt"
)

type person struct {
	name string
}
type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("I am", p.name)
}
func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		name: "Kenshi",
	}
	saySomething(&p1)
	//saySomething(p1)   	cannot use p1 (type person) as type human in argument to saySomething:
	//			person does not implement human (speak method has pointer receiver)
}
