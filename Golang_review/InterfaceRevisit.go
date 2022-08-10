// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math"
)

type square struct {
	lengh float32
}
type circle struct {
	radius float32
}

func (s square) area() float32 {
	a := s.lengh * s.lengh
	return a
}
func (s square) volumn() float32 {
	v := s.lengh * s.lengh * s.lengh
	return v
}
func (c circle) area() float32 {
	a := c.radius * c.radius * math.Pi
	return a
}

type shape interface {
	area() float32
}

func info(s shape) {
	fmt.Println(s.area())
}
func main() {

	s := square{
		5.4,
	}
	c := circle{
		4.3,
	}

	fmt.Println(s.area())
	fmt.Println(c.area())

	fmt.Println("volumn of s", s.volumn())

	info(s)
	info(c)
}
