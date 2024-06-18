package main

import "fmt"

func main() {
	var v, r float64
	const pi float64 = 3.14

	fmt.Scan(&r)
	v = (4.0 / 3.0) * pi * (r * r)
	fmt.Println(v)
}
