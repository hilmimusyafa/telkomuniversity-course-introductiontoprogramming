package main

import "fmt"

func main() {
	var x, h, i int

	fmt.Scan(&x)
	for h = 1; h <= x; h++{
		for i = 1 ; i <=x; i++ {
			// if h == 1 || h == x || i == 1 || i == x {
			// 	fmt.Print(i, " ")
			// } else {
			// 	fmt.Print("  ")
			// }
			fmt.Print(i, " ")
		}
		fmt.Println()
	}

}