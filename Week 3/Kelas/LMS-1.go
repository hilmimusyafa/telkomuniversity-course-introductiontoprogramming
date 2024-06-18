package main

import "fmt"

func main() {
	var r, L float64
	const pi float64 = 3.14

	fmt.Scan(&r)
	L = (pi * r * r)
	fmt.Println(L)
}
