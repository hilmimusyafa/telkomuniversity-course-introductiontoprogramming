package main

import "fmt"

func main() {
	var x, y float64
	var z bool
	//var stop bool

	//stop = false
	fmt.Scan(&x)
	for x <= 0 {
		y = y + x
		fmt.Scan(&x)
	}
	z = (y >= 10 || y == 10 )
	fmt.Print(z)
}