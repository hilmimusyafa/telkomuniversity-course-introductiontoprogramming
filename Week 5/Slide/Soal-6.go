package main

import "fmt"

func main() {
	var T, E, V int
	var stop, x bool

	stop = false
	fmt.Scan(&T)
	for !stop {
		fmt.Scan(&V)
		E += V
		//fmt.Println(E)
		x = E >= T
		fmt.Println(x)
		stop = (E >= T)
	}
}
