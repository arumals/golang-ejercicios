package main

import "fmt"

func main() {
	var marcasDeCoches = make([]string, 2)
	marcasDeCoches[0] = "Mazda"
	marcasDeCoches[1] = "Toyota"
	fmt.Println(marcasDeCoches)
	nuevoSlice := append(marcasDeCoches, "Nissan")
	fmt.Println(nuevoSlice)
}
