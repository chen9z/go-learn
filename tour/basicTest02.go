package main

import (
	"fmt"
	"strings"
)

func main() {
	sliceGameTest()
}

func func01() {
	i, j := 42, 2701
	p := &i
	println(*p)
	*p = 21
	println(i)

	p = &j
	*p = *p / 37
	println(j)
}

type Vertex struct {
	X int
	Y int
}

func func02() {
	fmt.Println(Vertex{1, 2})
}

func slice01() {
	primes := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(primes[1:4])
}

func slice02() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)

	s = nil
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func makeSlice01() {
	a := make([]int, 5)
	printSlice(a)

	a = make([]int, 0, 10)
	printSlice(a)
}

func sliceGameTest() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
