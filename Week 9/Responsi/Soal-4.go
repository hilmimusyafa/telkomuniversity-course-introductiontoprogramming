package main

import "fmt"

func main() {
	var m, s bool
	var l int
	var t, dt float64

	fmt.Scan(&m, &s, &l, &t)
	if m == true {
		if l < 8 {
			dt = t * 10 / 100
			t = t -dt
			fmt.Print(t)
		} else if l > 8{
			if s == true {
				dt = t * 20 / 100
				t = t - dt
				fmt.Print(t)
			} else {
				dt = t * 15 / 100
				t = t - dt
				fmt.Print(t)
			}
		} else if l > 12{
			if s == true {
				dt = t * 35 / 100
				t = t - dt
				fmt.Print(t)
			} else {
				dt = t * 30 / 100
				t = t - dt
				fmt.Print(t)
			}
		}
	} else if m == false {
		fmt.Print(t)
	}
}