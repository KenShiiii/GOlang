package main

import (
	"fmt"
)

func main() {
	defer f1()
	f2()
}

func f1() {
	fmt.Println("f1 called")
}
func f2() {
	fmt.Println("f2 called")
}

// Because of keyword defer, f1 is called after exiting main.
