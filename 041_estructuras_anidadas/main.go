package main

import "fmt"

type Ciudad struct {
	Nombre    string
	Poblacion int
}

type Pais struct {
	Nombre    string
	Capital   Ciudad
	Idioma    string
	Poblacion int
}

func main() {
	ciudadDeMexico := Ciudad{
		Nombre:    "Ciudad de Mexico",
		Poblacion: 20e6,
	}
	mexico := Pais{
		Nombre:    "México",
		Capital:   ciudadDeMexico,
		Idioma:    "Español",
		Poblacion: 120e6,
	}
	fmt.Printf("%+v\n", mexico)
}
