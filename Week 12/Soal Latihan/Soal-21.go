package main

import "fmt"

func main() {
	var x, y, a, b int

	fmt.Scan(&x, &y, &a, &b)
	x = x - 1
	for x < y{
		if (x % 10) + 1 == b {
			x = x + 1
		}
		if (x + 1) / 100 == a {
			x = x + 100
		}
		x = x + 1
		fmt.Println(x)
	}
}