package main

import "fmt"

func main() {
	var a, b, c, d int
	var x, y, z int

	fmt.Scan(&a, &b, &c)
	//d = (a+1) * (2*b + 2) * (3*c + 3)
	x = (a + 1)
	y = (2*b + 2)
	z = (3*c + 3)
	d = x * y * z
	fmt.Println(d)
}
