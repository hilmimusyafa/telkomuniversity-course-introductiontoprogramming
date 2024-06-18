package main

import "fmt"

func main() {
	var x, y, z int
	var h bool

	fmt.Scan(&x, &y, &z)
	h = (x+y < z)
	fmt.Print(h)
}
