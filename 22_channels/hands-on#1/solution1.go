package main

import (
	"fmt"
)
//solution #1 using anonymous func
func main() {
	c := make(chan int)
	go func(){
		c <- 42
	}()

	fmt.Println(<-c)
}
