package main

import (
	"fmt"
)

func div(n int, d int) bool {
	if n%d == 0 {
		return true
	}
	return false
}

func main() {

	for i := 1; i < 100; i++ {

		if div(i, 3) && div(i, 5) {
			fmt.Println("FIZZBUZZ")
		} else if div(i, 5) {
			fmt.Println("BUZZ")
		} else if div(i, 3) {
			fmt.Println("FIZZ")
		} else {
			fmt.Println(i)
		}

	}

	return
}
