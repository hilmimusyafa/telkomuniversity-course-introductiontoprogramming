package main

import "fmt"

func main() {
	var x, d1, d2, d3 int

	fmt.Scan(&x)
	if x < 0 {
		fmt.Println("Maaf angka input tidak boleh kurang dari 0, silahkan masukkan input lebih dari sama dengan 0")
		return
	} else if x > 999 {
		fmt.Println("Maaf angka input tidak boleh melebihi 999, silahkan masukkan input kurang dari sama dengan 999")
		return
	}

	d1 = (x / 100)
	d2 = (x / 10) % 10
	d3 = x % 10

	fmt.Println(d1, d2, d3)
}
