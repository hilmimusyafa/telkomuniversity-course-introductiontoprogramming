package main

import "fmt"

func main() {
	var j int
	var x, d int
	var stop bool

	fmt.Scan(&x)
	stop = false
	for !stop {
		d = x % 10
		x = x / 10
		j = j + d
		fmt.Print(d," ")
		stop = (x == 0)
	}
	fmt.Print(j)
}
