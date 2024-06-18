package main

import "fmt"

func main() {
	var h bool
	var n, m int
	var x, y, z int

	fmt.Scan(&n, &m)
	x = m % 10
	m = m / 10
	y = m % 10
	z = z + 1
	for y > x {
		x = m % 10
		m = m / 10
		y = m % 10
		z = z + 1
	}
	h = n == z
	fmt.Println(h)
}