package main

import "fmt"

func main() {
	var iter int
	var h int
	var f int

	fmt.Scan(&h)
	f = 1
	for iter = h; iter >= 1; iter-- {
		f *= iter
	}
	fmt.Println(f)
}
