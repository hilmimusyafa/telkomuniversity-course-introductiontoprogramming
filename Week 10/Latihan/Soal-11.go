package main

import "fmt"

func main() {
    var j1, j2, j3 string

	fmt.Scan(&j1, &j2, &j3)
	if j1 == "yes" {
		if j2 == "yes" {
			if j3 == "yes" {
				fmt.Print("lolos")	
			} else if j3 == "no" {
				fmt.Print("lolos")	
			}
		} else if j2 == "no"{
			if j3 == "yes" {
				fmt.Print("lolos")	
			} else if j3 == "no" {
				fmt.Print("tidak lolos")	
			}
		}
	} else if j1 == "no" {
		if j2 == "yes" {
			if j3 == "yes" {
				fmt.Print("lolos")	
			} else if j3 == "no" {
				fmt.Print("tidak lolos")	
			}
		} else if j2 == "no"{
			if j3 == "yes" {
				fmt.Print("tidak lolos")	
			} else if j3 == "no" {
				fmt.Print("tidak lolos")	
			}
		}
	}
}