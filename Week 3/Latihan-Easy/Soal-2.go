package main

import "fmt"

func main() {
	var k, r float64
	const pi float64 = 3.14

	fmt.Scan(&r)
	k = 2 * pi * r
	fmt.Println(k)
}
