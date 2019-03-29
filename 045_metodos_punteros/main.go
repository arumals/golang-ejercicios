package main

import "fmt"

type Banda struct {
	Nombre string
}

func (banda Banda) setNombre(nombre string) {
	banda.Nombre = nombre
}

func main() {
	banda1 := Banda{}
	banda1.setNombre("Metallica")
	fmt.Printf("%+v", banda1)
}
