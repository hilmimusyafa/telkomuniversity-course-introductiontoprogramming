package main

import "fmt"

func main() {
    var a, x, c int
	var k bool

	fmt.Scan(&a)
	x = 1 
	k = false
	for x <= a {
		if a % x == 0 {
			c = c + 1
		} 
		x = x + 1
	}
	if c == 2 {
		k = true
	}
	fmt.Print(k)	
}