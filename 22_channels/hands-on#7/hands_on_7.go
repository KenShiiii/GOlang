package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	c := make(chan int)

	go send(c)

	for v := range c {

		fmt.Println(v, "\tNums of Go routines: ", runtime.NumGoroutine())

	}
	fmt.Println("Done")

}
func send(c chan<- int) {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			for j := 1; j < 11; j++ {
				c <- i * j
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	close(c)

}
