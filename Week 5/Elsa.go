package main

import "fmt"

func main() {
	var x string

	fmt.Scan(&x)
	for x == "belum" {
		fmt.Println("cari beasiswa")
		fmt.Scan(&x)
	}
	fmt.Print("penacarian selesai")
}