package main

import (
	"fmt"
	"math"
)

type Length interface {
	len() float64
}

type geometry interface {
	area() float64
}

// a struct must implement all methods of interface to implement an interface
type circle struct {
	rad float64
}

func (c circle) area() float64 {
	return math.Pi * c.rad * c.rad
}

func (c circle) len() float64 {
	return 2 * math.Pi * c.rad
}

func measure(g geometry) {
	fmt.Println("Area of circle:", g.area())
}

func main() {
	c := circle{5}
	measure(c)

	var c1 Length = c
	println(c1.len())
}
