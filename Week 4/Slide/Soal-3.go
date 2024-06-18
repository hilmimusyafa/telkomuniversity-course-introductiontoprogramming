package main

import "fmt"

func main() {
	var iter int
	var h, j, p int
	var r float64

	fmt.Scan(&h)
	for iter = 1; iter <= h; iter++ {
		fmt.Scan(&j)
		p += j
	}
	r = float64(p) / float64(h)
	fmt.Println(r)
}
