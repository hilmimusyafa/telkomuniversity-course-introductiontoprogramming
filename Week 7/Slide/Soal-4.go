package main

import "fmt"

func main() {
	var diskon, cashback, kartu, ketersediaan bool
	var t int

	fmt.Scan(&t, &ketersediaan)
	kartu = (ketersediaan == true)
	diskon = (t >= 100000)
	cashback = (t >= 200000 && kartu == true)

	fmt.Println("Kartu?", kartu)
	fmt.Println("Diskon?", diskon)
	fmt.Println("Cashback?", cashback)
}
