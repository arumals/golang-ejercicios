package main

import "fmt"

func calcular(calculo func() float64) float64 {
	return calculo()
}

func main() {

	radio := 8.0

	area := func() float64 {
		return 3.1416 * radio * radio
	}

	perimetro := func() float64 {
		return 2 * 3.1416 * radio
	}

	fmt.Println("El radio del circulo es: ", radio)
	fmt.Println("El area es: ", calcular(area))
	fmt.Println("El perimetro es: ", calcular(perimetro))
}
