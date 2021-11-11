package main

import (
	"fmt"
	"runtime"
)

func main(){

	fmt.Println("Operating System:\t", runtime.GOOS)
	fmt.Println("Architecture:\t\t", runtime.GOARCH)
	fmt.Println("GO path:\t\t", runtime.GOROOT())

}
