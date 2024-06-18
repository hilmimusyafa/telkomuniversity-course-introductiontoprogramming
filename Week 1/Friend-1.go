package main

import "fmt"

func main() {
	var r, t int
	var v float64
	const pi float64 = 3.14

	fmt.Scan(&r, &t)
	v = 1.0 / 3.0 * pi * float64(r) * float64(r) * float64(t)
	fmt.Println(v)
}
