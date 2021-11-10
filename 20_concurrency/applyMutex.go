package main

import (
	"fmt"
	"runtime"
	"sync"
)


var counter int
var wg sync.WaitGroup
var mu sync.Mutex


func incrementer(){
	mu.Lock()
	i := counter
	runtime.Gosched()
	i++
	counter = i
	fmt.Println("counter: ",counter,"NumOfGoRountine: ",runtime.NumGoroutine())
	mu.Unlock()
	wg.Done()
}

func main(){
	counter = 0
	const gs = 100
	wg.Add(gs)
	for j := 0; j<gs;j++{
		go incrementer()

	}
	wg.Wait()
	fmt.Println("Final counter value: ",counter)

}
