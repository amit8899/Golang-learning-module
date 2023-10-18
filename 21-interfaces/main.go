package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
}

type circle struct {
	rad float64
}

func (c circle) area() float64 {
	return math.Pi * c.rad * c.rad
}

func measure(g geometry) {
	fmt.Println("Area of circle:", g.area())
}

func main() {
	c := circle{5}
	measure(c)
}
