package main

import "fmt"

const Pi = 3.1416

func circulo(radio float64) (area float64, perimetro float64) {
	area = Pi * radio * radio
	perimetro = 2 * Pi * radio
	return area, perimetro
}

func main() {
	a, p := circulo(8)
	fmt.Println("El area del circulo es: ", a)
	fmt.Println("El perimetro del circulo es: ", p)
}
