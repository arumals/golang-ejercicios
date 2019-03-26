package main

import "fmt"

type Auto struct {
	Marca    string
	Submarca string
	Modelo   int
	Color    string
}

func main() {
	auto1 := Auto{
		Marca:    "Toyota",
		Submarca: "Prius",
		Modelo:   2015,
		Color:    "blanco",
	}
	auto2 := Auto{
		Marca:    "Toyota",
		Submarca: "Corolla",
		Modelo:   2017,
		Color:    "cafe",
	}
	auto3 := Auto{
		Marca:    "Toyota",
		Submarca: "Prius",
		Modelo:   2015,
		Color:    "blanco",
	}
	if auto1 != auto2 {
		fmt.Println("Auto 1 y 2 son diferentes")
	}
	if auto1 == auto3 {
		fmt.Println("Auto 1 y 3 son iguales")
	}
}
