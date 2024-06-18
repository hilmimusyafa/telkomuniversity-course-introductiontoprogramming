package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	var j int

	fmt.Scan(&str)
	j = strings.Count(str, "go")
	fmt.Println(j)
}