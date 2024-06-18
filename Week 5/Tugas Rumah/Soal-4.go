package main

import "fmt"

func main() {
	var a, b int
	var jam, menit int
	var jj, jm int
	var stop bool

	stop = false
	for !stop {
		fmt.Scan(&a, &b)
		jam += a
		menit += b
		stop = (a == 0 && b == 0)
		//fmt.Print(jam, menit)
	}
	x := menit / 60
	jm = menit % 60
	jj = x + jam
	fmt.Print(jj, " ", jm)
}
