package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X float64
	Y float64
}

// Go does not have classes. However, one can define methods on types
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vertex) Output() string {
	return fmt.Sprintf("X: %v, Y: %v", v.X, +v.Y)
}
