package main

import "fmt"

type Pais struct {
	Nombre    string
	Capital   string
	Idioma    string
	Poblacion int
}

func main() {
	irlanda := Pais{
		Nombre:    "Irlanda",
		Capital:   "Dublin",
		Idioma:    "Irlandes",
		Poblacion: 4857000,
	}
	fmt.Printf("%+v\n", irlanda)
}
