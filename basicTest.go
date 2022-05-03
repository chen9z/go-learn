package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	deferTest()
}

func switchForOs() {
	fmt.Println("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Println("%S", os)
	}
}

func switchWeekDay() {
	fmt.Println("when's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today:
		fmt.Println("today")
	case today + 1:
		fmt.Println("tomorrow")
	case today + 2:
		fmt.Println("in two days")
	default:
		fmt.Println("too far wasy")
	}
}

func switchNoCondition() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		println("Good morning")
	case t.Hour() < 17:
		println("Good afternoon")
	default:
		println("Good evening")
	}
}

func deferTest() {
	println("counting")
	for i := 0; i < 10; i++ {
		defer println(i)
	}
	println("done")
}
