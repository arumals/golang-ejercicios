package main

import "fmt"

type Pais struct {
	Nombre    string
	Capital   string
	Idioma    string
	Poblacion int
}

func main() {
	colombia := new(Pais)
	fmt.Printf("%+v\n", colombia)
	colombia.Nombre = "Colombia"
	colombia.Capital = "Bogota"
	colombia.Idioma = "Español"
	colombia.Poblacion = 49e6
	fmt.Printf("%+v\n", colombia)
}
