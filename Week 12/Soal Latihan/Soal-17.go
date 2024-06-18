package main

import "fmt"

func main() {
	var a1, a2, a3 int
	var t int
	
	fmt.Scan(&a1, &a2, &a3)
	if a1 <= 6 && a2 <= 6 && a3 <= 6 {
		for a1 + a2 + a3 < 15{
			if (a1 % 2 == 1 && a2 % 2 == 1 && a3 % 2 == 1) || (a1 % 2 == 0 && a2 % 2 == 0 && a3 % 2 == 0){
				t = t + 1
			}
			fmt.Scan(&a1, &a2, &a3)
		}
	}	
	fmt.Print(t)
}