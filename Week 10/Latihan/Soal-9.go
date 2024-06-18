package main

import "fmt"

func main() {
    var c string

	fmt.Scan(&c)
	if c == "tinggi" {
		fmt.Print("macet")
	} else {
		fmt.Print("tidak macet")
	}
}