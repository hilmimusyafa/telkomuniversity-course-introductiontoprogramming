package main

import "fmt"

func main() {
	var t int
	var k bool

	fmt.Scan(&t)
	k = false
	if (t % 4 == 0  && t % 100 != 0 || t % 400 == 0) {
		k = true
	}
	fmt.Print(k)
}