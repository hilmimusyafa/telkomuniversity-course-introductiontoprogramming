package main

import "fmt"

func main() {
	var x, iter int

	for iter = 1 ; iter <= 100; iter++ {
		x = x + iter
	}
	fmt.Print(x)
}