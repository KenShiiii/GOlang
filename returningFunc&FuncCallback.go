package main

import (
	"fmt"
)

func main() {
	f := func() {
		fmt.Println("what up! world")
	}
	g := callback(f)
	g()
}
//A “callback” is when we pass a func into a func as an argument
func callback(f func()) func() {
	return f
}
