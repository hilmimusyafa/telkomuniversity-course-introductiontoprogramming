package main

import "fmt"

func main() {
	var h, s int
	var selisih int

	fmt.Scan(&h, &s)
	selisih = h - s

	if selisih != 0 {
		if (selisih % 2 == 0) {
			if (selisih <= -16) && (selisih > 0) {
				fmt.Print("Shinchi menang")
			} else {
				fmt.Print("Shinichi dan Heiji kalah")	
			}
		} else if (selisih % 2 == 1) {
			if (selisih > 0) && (selisih <= 45) {
				fmt.Print("Heiji menang")
			} else {
				fmt.Print("Shinichi dan Heiji kalah")	
			}
	} else if selisih == 0 {
		fmt.Print("Shinchi dan Heiji seri")
	} else {
		fmt.Print("Shinchi dan Heiji kalah")}
	}
}