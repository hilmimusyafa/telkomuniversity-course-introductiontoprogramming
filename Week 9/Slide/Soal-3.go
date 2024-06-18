package main

import "fmt"

func main() {
	var a, b, c int
	var max, mid, min int

	fmt.Scan(&a, &b, &c)
	max = a
	if b > a && b > c{
		max = b
	}else if c > a && c > b{
		max = c
	}
	mid = b
	if (a < b && a > c) || (c > a && b < a) {
		mid = a
	}else if (c < b && c > a) || (a > c && b < c) {
		mid = c
	}
	min = c
	if b < a && b < c{
		min = b
	}else if a < b && a < c{
		min = a
	}
	fmt.Print(min, mid, max)
}