package main

import "fmt"

func main () {
	var iter, n, d, jumlah int
	var rata float64
	
	fmt.Scan(&n)
	for iter = 1; iter <= n ; iter++ {
		fmt.Scan(&d)
		jumlah += d
	}
	rata = float64(jumlah) / float64(n)
	fmt.Println(jumlah)
	fmt.Printf("%.2f\n", rata)
}