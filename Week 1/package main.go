package main

import "fmt"

func main() {
	// Menampilkan teks ke layar
	fmt.Println("Hello, World!")

	// Variabel dan penggunaan if-else
	x := 10
	if x > 5 {
		fmt.Println("x lebih besar dari 5")
	} else {
		fmt.Println("x tidak lebih besar dari 5")
	}

	// Perulangan for
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Slice (potongan array)
	fruits := []string{"apel", "pisang", "jeruk"}
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}
}
