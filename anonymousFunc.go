package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("anonymous func ran")
	}()

	func(i int) {
		fmt.Println(i, "times 8 equals", i*8)
	}(34)

}
