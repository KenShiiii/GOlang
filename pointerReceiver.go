package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person) {
	p.age++
	//(*p).age++         // work as well

}
func main() {
	me := person{
		first: "Kenshi",
		last:  "Kuo",
		age:   25,
	}
	fmt.Println(me)
	changeMe(&me)
	fmt.Println(me)

}
