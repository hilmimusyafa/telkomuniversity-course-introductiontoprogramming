package main

import "fmt"

func main() {
	var a, b, c int
	var x bool
	//var stop bool

	fmt.Scan(&a)
	//stop = false
	x = true
	for a >= 10 {
		b = a % 10
		a /= 10
		c = a % 10
		x = x && (b-c == 1 || c-b == 1)
	}
	fmt.Print(x)
}
