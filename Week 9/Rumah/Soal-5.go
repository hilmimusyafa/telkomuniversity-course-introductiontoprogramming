package main

import "fmt"

func main() {
	var x1, x2, x3 int
	var min int

	fmt.Scan(&x1, &x2, &x3)
	
	min = x3
	if x2 < x1 && x2 < x3{
		min = x2
	}else if x1 < x2 && x1 < x3{
		min = x1
	}
	fmt.Print(min)
}