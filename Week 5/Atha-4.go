package main

import "fmt"

func main() {
	var x, hasil float64
	
	fmt.Scan(&x)
	hasil = 1
	for x != 0 {
		y := (1 / x)
		hasil = y * hasil
		fmt.Scan(&x)
	}
	fmt.Printf("%.5f", hasil)
}