package main

import "fmt"

func main() {
	var f, w, x, y, z float64

	fmt.Scan(&w, &x, &y, &z)
	f = (1 + 3*w*w) / (4*x - 5*y*y*y + 6*z*z*z*z)
	fmt.Println(f)
}
