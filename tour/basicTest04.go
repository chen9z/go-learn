package main

import (
	"fmt"
	"math"
)

func main() {
	f := myFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

type Abser interface {
	Abs() float64
}

type myFloat float64

func (f myFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
