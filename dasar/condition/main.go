package main

import (
	"fmt"
)

func main() {
	var jk = "male"
	// Kondisi If
	if true {
		fmt.Println("True")
	}

	// Kondsi If dengan Else
	if false {
		fmt.Println("Bukan False")
	} else {
		fmt.Println("True")
	}

	//ternary operator golang
	a := (map[bool]string{true: "L", false: "p"})[jk == "male"]
	fmt.Println(a)

	variableAngka := 2
	// Kondisi If berulang / nested if
	if false {
		fmt.Println("Angka = 1")
	} else if true {
		fmt.Println("Angka Bukan 1, Angkanya = 2")
	} else {
		fmt.Println("Angka Bukan 1 & 2")
	}

	// Kondisi menggunakan switch
	switch variableAngka {
	case 1:
		fmt.Println("Angka = 1")
	case 2:
		fmt.Println("Angka Bukan 1, Angka nya = 2")
	default:
		fmt.Println("Angka Bukan 1 & 2")
	}
}
