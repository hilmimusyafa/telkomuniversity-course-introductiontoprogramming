package main

import "fmt"

func main() {
    var h, p string

	fmt.Scan(&h, &p)
	if h == "ya" {
		if p == "ya" {
			fmt.Print("berangkat")
		} else if p == "tidak" {
			fmt.Print("diam di rumah")
		}
	} else if h == "tidak" {
		fmt.Print("berangkat")
	}
}