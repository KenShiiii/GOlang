package main

import (
	"fmt"
)

func main() {
	f1 := func() {
		fmt.Println("anonymous func ran")
	}

	f2 := func(i int) {
		fmt.Println(i, "times 8 equals", i*8)
	}

	f1()
	f2(46)
}
