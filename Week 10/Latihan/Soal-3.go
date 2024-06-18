package main

import "fmt"

func main() {
	var m1, m2, m3 string

	fmt.Scan(&m1, &m2, &m3)
	if m1 == m2 && m1 == m2 && m2 == m3 {
		fmt.Print("imbang")
	} else {
		if m1 == m2 {
			fmt.Print("pemain 3 pemenang")
		} else if m1 == m3 {
			fmt.Print("pemain 2 pemenang")
		} else if m2 == m3 {
			fmt.Print("pemain 1 pemenang")
		}
	}
}
