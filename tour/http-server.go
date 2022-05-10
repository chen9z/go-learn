package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println(w, "hello")
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe("localhost:8090", nil)
}
