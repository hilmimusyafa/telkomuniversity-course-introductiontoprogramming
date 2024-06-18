package main

import "fmt"

func main() {
	var L, r float64
	const pi float64 = 22.0 / 7.0

	fmt.Scan(&r)
	//pi = 22.0 / 7.0
	L = 4 * pi * (r * r)
	fmt.Println(L)
}
