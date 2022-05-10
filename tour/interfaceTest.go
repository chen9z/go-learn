package main

import (
	"fmt"
	"math"
)

func main() {
	var k interface{}
	describeI(k)

	var i I
	describe(i)
	i = &T{"hello"}
	describe(i)

	i = F(math.Pi)
	describe(i)

	i.M()
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
func describeI(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
