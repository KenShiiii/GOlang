package main

import (
	"fmt"
	"runtime"
	"sync"
)


var n int
var wg sync.WaitGroup

func incrementer(){
	wg.Add(1)
	i := n
	runtime.Gosched()
	i++
	n = i

	wg.Done()
}

func main(){
	n = 0
	for j := 0; j<100;j++{
		go incrementer()
		fmt.Println("n: ",n,"NumOfGoRountine: ",runtime.NumGoroutine())


	}
	wg.Wait()

}
