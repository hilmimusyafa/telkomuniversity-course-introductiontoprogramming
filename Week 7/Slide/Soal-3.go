package main

import "fmt"

func main() {
	var a, b, c int
	var k bool

	fmt.Scan(&a, &b, &c)
	k = (a+b > c && a+c > b && b+c > a)
	fmt.Print(k)
}
