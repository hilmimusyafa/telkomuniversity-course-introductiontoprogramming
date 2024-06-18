package main

import "fmt"

func main() {
	var m, t int
	var m1, m2, m3, m4, m5 int

	fmt.Scan(&m)
	if m == 4{
		fmt.Scan(&m1, &m2, &m3, &m4)
		t = m1 + m2 + m3 + m4
		fmt.Print(t)
	} else if m == 5 {
		fmt.Scan(&m1, &m2, &m3, &m4, &m5)
		t = m1 + m2 + m3 + m4 + m5
		fmt.Print(t)
	} else {
		fmt.Print("Minggu kurang/lebih")
	}
}