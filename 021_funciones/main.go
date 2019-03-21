package main

import "fmt"

const Pi = 3.1416

func area(radio float64) float64 {
	return Pi * radio * radio
}

func main() {
	fmt.Println("El area de un circulo cuyo radio es 3 es: ", area(3))
}
