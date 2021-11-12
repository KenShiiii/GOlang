package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	for v := range c{
		fmt.Println(v)

	}
	fmt.Println("Done")

	/*			This was my original answer, but using range is much cleaner and easier
	func() {
		for {
			v, ok := <-c
			if ok {

				fmt.Println(v)
			} else {
				fmt.Println("channel c is closed")
				return
			}

		}
	}()
	*/
}
