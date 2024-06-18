package main

import "fmt"

func main() {
	var n int
	var a, b, c, d, e, x bool
	var i int

	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		fmt.Scan(&a, &b, &c, &d, &e)
		x = x && (a == true && b == true && c == true && d == true && e == true)
	}
	fmt.Println(x)
}
