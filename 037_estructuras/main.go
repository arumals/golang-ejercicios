package main

import "fmt"

type Pais struct {
	Nombre     string
	Capital    string
	Idioma     string
	Habitantes int
}

func main() {
	irlanda := Pais{
		Nombre:     "Irlanda",
		Capital:    "Dublin",
		Idioma:     "Irlandes",
		Habitantes: 4857000,
	}
	fmt.Println(irlanda)
}
