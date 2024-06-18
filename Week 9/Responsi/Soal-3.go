package main

import "fmt"

func main() {
	var s1, s2, s3 int

	fmt.Scan(&s1, &s2, &s3)
	if s1 > 0 && s2 > 0 && s3 > 0 {
		if s1 == s2 || s1 == s3 || s2 == s3 {
			fmt.Print("Segitiga Sama Kaki")
		} else if s1 == s2 && s1 == s3 && s2 == s3 {
			fmt.Print("Segitiga Sama Sisi")
		} else if s1 * s1 == s2 + s3 || s2 * s2 == s1 + s3 || s3 * s3 == s1 + s2 {
			fmt.Print("Segitiga Siku - Siku")
		} else {
			fmt.Print("Segitiga Sembarang")
		}
	} else if s1 >=  s2 + s3 || s2 >= s1 + s3 || s3 >= s1 + s2 {
		fmt.Print("Segitiga tidak dapat dibentuk")
	} else if s1 <= 0 || s2 <= 0 || s3 <= 0{
		fmt.Print("Bukan Segitiga")
	}
}