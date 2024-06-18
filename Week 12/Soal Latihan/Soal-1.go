package main

import "fmt"

func main() {
    var a1, a2, t int
	
	fmt.Scan(&a1, &a2)
	if a1 % 2 != 0 && a2 % 2 == 0 {
		t = t + a1 + a2
	}
	for a1 != 0 && a2 != 0 {
		fmt.Scan(&a1, &a2)
		if a1 % 2 != 0 && a2 % 2 == 0 {
			t = t + a1 + a2
		}
	}
	fmt.Print(t)
}