package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)


var counter int64
var wg sync.WaitGroup



func incrementer(){

	atomic.AddInt64(&counter, 1)
	runtime.Gosched()

	fmt.Println("counter: ",atomic.LoadInt64(&counter),"\tNumOfGoRountine: ",runtime.NumGoroutine())

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
