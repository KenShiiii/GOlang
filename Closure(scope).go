package main

import (
	"fmt"
)

func main() {
	a := foo()
	b := foo()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

func foo() func() int {
	var x int
	return func() int {
		x++
		return x
	}

}
