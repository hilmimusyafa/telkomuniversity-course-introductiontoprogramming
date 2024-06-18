package main

import "fmt"

func main() {
	var d1, d2, k int
	var x bool

	fmt.Scan(&d1, &d2)
	if (d1 % 2 == 0 && d2 % 2 == 0) || (d1 % 2 == 1 && d2 % 2 == 1) {
		x = true
		k = 1
	}
	for x == false {
		fmt.Scan(&d1, &d2)
		if (d1 % 2 == 0 && d2 % 2 == 0) || (d1 % 2 == 1 && d2 % 2 == 1) {
			x = true
		}
		k = k + 1
	}
	fmt.Print(k)
}