package main

import "fmt"

func main() {
	var n, m, x, y, i float64
	var hasil float64

	fmt.Scan(&n)
	fmt.Scan(&m)
	hasil = 0
	for i = n; i <= m; i++ {
		y = 2 / (n + x)
		hasil = hasil + y
		x = x + 1
	}
	fmt.Printf("%.2f", hasil)
}
