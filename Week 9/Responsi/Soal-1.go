package main

import "fmt"

func main() {
	var k rune
	var x bool

	fmt.Scanf("%c", &k)
	x = false
	if (k >= 'a' && k <= 'z') || (k >= 'A' && k <= 'Z') || (k >= '0' && k <= '9') {
		x = true
	}
	fmt.Print(x)
}
