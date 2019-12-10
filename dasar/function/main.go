package main

import (
	"fmt"
)

// Function Biasa
func kaliDua(angka int) int {
	return angka * 2
}

func isTrue(isTrue int) bool {
	if isTrue == 1 {
		return true
	}
	return false
}

func PrintSum(name string, values ...int) {
	sum := 0
	// Loop over all variadic arguments and sum them.
	for i := range values {
		sum += values[i]
	}
	fmt.Println(name, sum)
}
func main() {
	fmt.Println("Fungsi return value integer", kaliDua(2))
	fmt.Println("function dengan nilai kembalian boolean dengan value", isTrue(0))
	PrintSum("cat", 1, 2)
	data := "Tampilkan Di Fungsi Literal"
	// Function Literal / Lambda Functin / Anonymous Functin
	func(dataArgumen string) {
		fmt.Println("Fungsi Literal", dataArgumen)
	}(data)

}
