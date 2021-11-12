package main

import (
	"fmt"
)
//solution #2 using buffered channel
func main() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
