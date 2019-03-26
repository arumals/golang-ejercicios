package main

import "fmt"

type Pais struct {
	Nombre    string
	Capital   string
	Idioma    string
	Poblacion int
}

func main() {
	panama := Pais{
		Nombre:    "Panama",
		Capital:   "Ciudad de Panama",
		Idioma:    "Español",
		Poblacion: 4170607,
	}
	fmt.Println("Nombre: ", panama.Nombre)
	fmt.Println("Capital: ", panama.Capital)
	fmt.Println("Idioma: ", panama.Idioma)
	fmt.Println("Población: ", panama.Poblacion)
}
