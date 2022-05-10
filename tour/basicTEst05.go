package main

import "fmt"

func main() {

	typeAssert01()

}

func typeAssert01() {

	var i interface{} = "hello"
	s := i.(string)

	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	//s, ok := i.(float64)
	//fmt.Println(s, ok)

}
