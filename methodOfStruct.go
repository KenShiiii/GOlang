<<<<<<< HEAD
package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last)
}

func main() {
	p1 := person{
		first: "Kenshi",
		last:  "Kuo",
		age:   25,
	}
	p1.speak()
}
=======
package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last)
}

func main() {
	p1 := person{
		first: "Kenshi",
		last:  "Kuo",
		age:   25,
	}
	p1.speak()
}
>>>>>>> 242ba4965e8490bc4fdc757e67a40234c255ac89
