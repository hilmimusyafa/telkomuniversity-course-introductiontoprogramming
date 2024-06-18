package main

import "fmt"

func main() {
	var i, j float64
	var stop bool

	j = 0
	fmt.Scan(&i)
	stop = false
	for !stop {
		j += i
		fmt.Scan(&i)
		stop = (i == 0)
	}
	fmt.Print(j)
}
