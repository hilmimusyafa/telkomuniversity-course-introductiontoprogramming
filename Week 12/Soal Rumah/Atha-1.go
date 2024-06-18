package main

import (
	"fmt"
	"strconv"
)

func main() {
	var kata string
	fmt.Scan(&kata)

	totalAngka := totalKarakter(kata)
	fmt.Println(totalAngka)
}

func totalKarakter(kata string) int {
	var totalAngka int

	for i := 0; i < len(kata); i++ {
		if isTotalKarakter(kata[i]) {
			num, err := strconv.Atoi(string(kata[i]))
			if err == nil {
				totalAngka += num
			}
		}
	}

	return totalAngka
}

func isTotalKarakter(char byte) bool {
	return char >= '0' && char <= '9'
}