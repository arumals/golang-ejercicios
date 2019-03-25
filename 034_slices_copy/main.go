package main

import "fmt"

func main() {
	figuras := []string{"circulo", "cuadrado", "triangulo", "rombo", "trapecio", "heptagono"}
	var figuras2 = make([]string, len(figuras))
	fmt.Println(figuras)
	copy(figuras2, figuras)
	figuras = append(figuras[:1], figuras[2:]...)
	fmt.Println(figuras2)
	fmt.Println(figuras)
}
