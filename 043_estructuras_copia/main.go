package main

import "fmt"

type Pais struct {
	Nombre string
}

func main() {
	pais1 := Pais{Nombre: "Canada"}
	pais2 := &pais1
	pais1.Nombre = "Inglaterra"
	fmt.Printf("%+v\n", pais1)
	fmt.Printf("%+v\n", *pais2)
}
