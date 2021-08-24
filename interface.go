package main

import (
	"fmt"
	"math"
)

type square struct {
	long  float32
}
type circle struct {
	radius float32
}
type shape interface {
	area() float32
}

func (s square) area() float32 {
	a := s.long * s.long
	return a
}
func (c circle) area() float32 {
	a := math.Pi * c.radius * c.radius
	return a
}
func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	sq := square{
		long:  15.5,
	}
	cc := circle{
		radius: 5.0,
	}
	//fmt.Println(sq.area())
	//fmt.Println(cc.area())
	info(sq)
	info(cc)
}
