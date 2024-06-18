package main

import "fmt"

func main() {
	var js, jA, jB, p int

	fmt.Scan(&js, &jA, &jB)
	p = 60 / 100 * js
	if jA+jB < p {
		fmt.Print("Tidak ada pemenang")
	} else if jA+jB > p {
		if jA > jB {
			fmt.Print("Kandidat A menang")
		} else {
			fmt.Print("Kandidat B menang")
		}
	}
}
