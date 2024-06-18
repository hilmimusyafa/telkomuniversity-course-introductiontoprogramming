package main

import "fmt"

func main() {
	var a, b, c float64
	var x bool

	fmt.Scan(&a, &b, &c)
	x = (a+b > c && a+c > b && b+c > a)
	// (a>0 && b>0 && c>0)
	fmt.Println(x)
}
