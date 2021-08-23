package main

import (
	"fmt"
)

func main() {
	x := 25
	fmt.Println(&x)
	fmt.Printf("%T\n", &x)
	fmt.Println(*&x)
	fmt.Printf("%T\n", *&x)
	y := &x
	fmt.Println(*y)

}
