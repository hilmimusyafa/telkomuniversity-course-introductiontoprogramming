package main

import "fmt"

func main() {
	var a, b, c int
	var x bool

	fmt.Scan(&a, &b, &c)
	x = (float64(a) + float64(b)) / 2 == float64(c) || (float64(a) + float64(c)) / 2 == float64(b) || (float64(b) + float64(c)) / 2 == float64(a)
	fmt.Print(x)
}