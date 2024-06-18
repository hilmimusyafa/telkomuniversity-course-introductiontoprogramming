package main

import "fmt"

func main() {
	var d, at, m, o, w, x int
	//var w, x, int

	fmt.Scan(&d)
	w = d % 6000
	at = d / 6000

	x = w % 100
	m = w / 100

	o = x * 6
	fmt.Println(at, m, o)
}
