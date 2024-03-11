package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}
type square struct {
	length float64
}
type circle struct {
	radius float64
}

func (s *square) area() float64 {
	return s.length * s.length
}

func (c *circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func toString(s shape) string {
  return fmt.Sprint("area of the shape is: ", s.area())
}

func main() {
	shapes := []shape{
		&square{length: 10},
		&square{length: 2},
		&circle{radius: 2},
		&circle{radius: 10},
	}
	for _, v := range shapes {
    fmt.Println(toString(v))
	}
}
