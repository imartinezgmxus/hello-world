package main

import (
	"fmt"
)

func main() {
	var x string
	x = "¡Hola Mundo!"
	fmt.Println(x)
	ConvertPlus()
	ForExample()
}

// ConvertPlus - Convert a string variable to Float64
func ConvertPlus() {
	a := 7
	b := 14.0

	fmt.Printf("Tipo de variable para A %T \n", a)
	fmt.Printf("Tipo de variable para B %T \n", b)

	fmt.Println("Resultado de la suma: ", float64(a)+b)
}

// ForExample - For ForExample
func ForExample() {
	i := 0
	for i <= 9 {
		fmt.Println(i)
		i++
	}
}
