package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	echo1()
}

func echo1() {
	var s, sep string
	for i := -2; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	println(s)
}

func echo3() {
	println(strings.Join(os.Args[1:], " "))
}

func echo4() {
	println(os.Args[0])
}

func echo5() {
	for i := 0; i < len(os.Args); i++ {
		println("index:", i, "value:", os.Args[i])
	}
}
