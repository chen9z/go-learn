package main

import (
	"fmt"
	"math"
)

func main() {
	closures()
}

func rangeTest01() {
	var ss = []string{"1", "2", "3"}
	for i, v := range ss {
		fmt.Printf("%d,%s\n", i, v)
	}
}

func rangTest02() {
	var ss = []string{"1", "2", "3"}
	for i := range ss {
		println(i)
	}
}

func testMap01() {
	m := make(map[string]string)
	m["a"] = "a"
	m["b"] = "b"
	delete(m, "a")

	fmt.Println(m)
}

func wordCount(s string) map[string]int {
	return map[string]int{"x": 1}
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func funValue() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func closures() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}
