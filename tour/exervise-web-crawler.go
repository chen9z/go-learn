package main

import "fmt"

func main() {
	var i = 0
	fmt.Println(i)
}

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher) {

}
