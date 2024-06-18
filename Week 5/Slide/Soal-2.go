package main

import "fmt"

func main() {
	var j float64
	var km float64
	var stop bool

	stop = false
	for !stop {
		fmt.Scan(&km)
		stop = (km == 0)
		j = j + km
	}
	fmt.Println(j)
}
