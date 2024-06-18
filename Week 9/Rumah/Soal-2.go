package main

import "fmt"

func main() {
	var h int

	fmt.Scan(&h)
	if h >= 1 && h <= 7 {
		if h == 1 {
			fmt.Print("minggu")
		} else if h == 2 {
			fmt.Print("senin")
		} else if h == 3 {
			fmt.Print("selasa")
		} else if h == 4 {
			fmt.Print("rabu")
		} else if h == 5 {
			fmt.Print("kamis")
		} else if h == 6 {
			fmt.Print("jumat")
		} else if h == 7 {
			fmt.Print("sabtu")
		} 
	} else {
		fmt.Print("diluar range")
	}
} 