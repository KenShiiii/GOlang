package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}
func receive(c, q <-chan int) {

	for {
		select {
		case n, ok := <-c:
			if ok {
				fmt.Println("Received from channel c :", n)
			} else {
				fmt.Println("Channel c is closed")
				return
			}
		}

	}

}

func gen(q <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}
