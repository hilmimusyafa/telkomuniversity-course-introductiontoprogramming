package main

import "fmt"

func main() {
	var bmi, berat, tinggi float64
    
    fmt.Scan(&tinggi, &berat)
    
	bmi = berat / tinggi * tinggi
	if bmi < 18.5 {
		fmt.Print("Kurang Berat Badan")
	}else if 18.5 <= bmi && bmi <= 22.9 {
	    fmt.Print("Normal")
	}else {
	    fmt.Print("Berat Badan Lebih")
	}
}