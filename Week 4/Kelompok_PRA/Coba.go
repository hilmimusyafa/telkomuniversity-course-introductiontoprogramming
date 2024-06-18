package main

import "fmt"

func main() {
	var i, a, b, c, d, e, f int

	i = 1
	a = i 
	b = i + 1
	c = i + 2
	d = i + 3
	e = i + 4
	f = i + 5
	i = a, b, c, d, e, f
	fmt.Println(i)
}
