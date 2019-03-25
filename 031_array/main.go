package main

import "fmt"

func main() {
	var marcasDeAutos [4]string
	marcasDeAutos[0] = "Mazda"
	marcasDeAutos[1] = "Toyota"
	marcasDeAutos[2] = "Nissan"
	marcasDeAutos[3] = "Mitsubishi"
	fmt.Println(marcasDeAutos)
	fmt.Println("Marca 1: ", marcasDeAutos[0])
}
