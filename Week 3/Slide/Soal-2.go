package main

import "fmt"

func main() {
	var h, p, q, r, s, x int

	fmt.Scan(&x)
	//Pengambilan setiap digit
	p = (x / 10)
	q = x % 10
	// Pengalian 11
	r = p * 11
	s = q * 11
	// Hasil
	h = (r * 100) + s
	fmt.Println(h)
}
