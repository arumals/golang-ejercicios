package main

import "fmt"

type Pais struct {
	Nombre    string
	Capital   string
	Idioma    string
	Poblacion int
}

func main() {
	var belice Pais
	fmt.Printf("%+v\n", belice)
	panama := Pais{
		Nombre:    "Panama",
		Capital:   "Ciudad de Panama",
		Idioma:    "Español",
		Poblacion: 4170607,
	}
	fmt.Printf("%+v\n", panama)
}
