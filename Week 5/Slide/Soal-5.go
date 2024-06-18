package main

import "fmt"

func main() {
	var a, b, c, d int
	var x bool
	var stop bool

	fmt.Scan(&a)
	stop = false
	for !stop {
		b = a % 10
		c = a % 100
		d = c / 10
		a = a / 10
		x = (d == b+1 || d == b-1)
		stop = (b != b+1 || b != b-1)
	}
	fmt.Print(x)
}
