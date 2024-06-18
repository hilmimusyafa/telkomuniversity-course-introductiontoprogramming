package main

import "fmt"

func main() {
	var t, i, x, y, z int

	fmt.Scan(&x, &y, &z)
	for i = 2; i <= z; i++ {
		if i % 2 == 0{
			t = t + x
		} 
		if i % 3 == 0{
			t = t + x + y
		}
	}
	fmt.Print(t)
}