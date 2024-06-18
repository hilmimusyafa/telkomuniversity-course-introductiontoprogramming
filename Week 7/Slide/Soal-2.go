package main

import "fmt"

func main() {
	var t int
	var k bool

	fmt.Scan(&t)
	t = t % 4
	k = (t == 0)
	fmt.Print(k)
}
