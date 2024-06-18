package main

import "fmt"

func main() {
	var n, x, y, z, p int
	var tx, ty, tz int

	fmt.Scan(&n)
	for p < n {
		fmt.Scan(&x, &y, &z)
		tx = tx + x
		ty = ty + y
		tz = tz + z
		p = p + 1
		x = 0 
		y = 0
		z = 0
	}
	if tx > ty && tx > tz {
		fmt.Print("+", " ", tx)
	} else if ty > tx && ty > tz {
		fmt.Print("="," ", ty)
	} else {
		fmt.Print("-"," ", tz)
	}
}