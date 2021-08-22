package main

import (
	"fmt"
)

const pi = 3.141615

type square struct {
	long  float32
	width float32
}
type circle struct {
	radius float32
}
type shape interface {
	area() float32
}

func (s square) area() float32 {
	a := s.long * s.width
	return a
}
func (c circle) area() float32 {
	a := pi * c.radius * c.radius
	return a
}
func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	sq := square{
		long:  15.5,
		width: 13.0,
	}
	cc := circle{
		radius: 5.0,
	}

	//fmt.Println(sq.area())
	//fmt.Println(cc.area())
	info(sq)
	info(cc)
}
