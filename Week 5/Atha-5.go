package main

import "fmt"

func main() {
	var x, y int
	var a, b, h ,i, z int
	var stop bool

	stop = false
	fmt.Scan(&x)
	a = x / 10
	b = x % 10
	for !stop {
		fmt.Scan(&y)
		h = y / 10
		i = y % 10
		z = z + 1
		stop = !(a == h || a == i || b == h || b == i)
	}
	fmt.Print(z)
}