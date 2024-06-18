package main

import "fmt"

func main() {
    var k string
	var j, tbagus, tcacat int

	fmt.Scan(&k, &j)
	for k != "selesai" && j != 0 {
		if k == "bagus" {
			tbagus = tbagus + j
		} else if k == "cacat" {
			tcacat = tcacat + j
		}
		fmt.Scan(&k, &j)
	}
	fmt.Print(tbagus, tcacat)
}	