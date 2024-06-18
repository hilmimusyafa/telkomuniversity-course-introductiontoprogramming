package main

import "fmt"

func main() {
	var total, a int

	a = 1
	for a <= 1000 {
		total += a
		a++
	}
	fmt.Println(total)
}
