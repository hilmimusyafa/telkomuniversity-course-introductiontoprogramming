package main

import "fmt"

func main() {
	var n float64
	var r, j, b float64
	//var stop bool

	//stop = false
	fmt.Scan(&n)
	for n >= 0 {
		j = j + n
		b = b + 1
		fmt.Scan(&n)
	}
	r = j / b
	fmt.Printf("%.2f\n", r)
}
