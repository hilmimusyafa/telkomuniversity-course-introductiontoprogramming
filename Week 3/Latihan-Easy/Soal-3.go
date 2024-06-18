package main

import "fmt"

func main() {
	var l, r float64
	const pi float64 = 3.14

	fmt.Scan(&r)
	l = pi * r * r
	fmt.Println(l)
}
