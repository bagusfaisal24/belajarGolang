package main

import "fmt"

type Hellow interface {
	getSalam() string
}
type Count interface {
	hitung(a int) int
}

type botInggris struct{}

type botSpayol struct{}

type botHitung struct {
	a int
}

// tidak perlu deklarasi objeknya karena tidak digunakan func (i botInggris) ...
func (botInggris) getSalam() string {
	return "Halo"
}

func (botSpayol) getSalam() string {
	return "Hola"
}

func (number botHitung) hitung(a int) int {
	return number.a * a
}

func main() {
	objekInggris := botInggris{}
	objekSpanyol := botSpayol{}
	objectNumberHitung := botHitung{2}

	printSalam(objekInggris)
	printSalam(objekSpanyol)
	printHitung(objectNumberHitung)
}

func printSalam(hellow Hellow) {
	fmt.Println(hellow.getSalam())
}

func printHitung(count Count) {
	fmt.Println(count.hitung(5))
}
