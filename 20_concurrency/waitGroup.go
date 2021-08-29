package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("CPU: ", runtime.NumCPU())
	wg.Add(2)
	go func() {
		fmt.Println("this prints from a go routine.")
		wg.Done()
	}()
	go func() {
		fmt.Println("this prints from another go routine.")
		wg.Done()
	}()

	wg.Wait()

}
