package main

import "fmt"

func hola(nombre *string) {
	fmt.Println(*nombre, nombre)
}

func main() {
	javier := "Javier"
	fmt.Println(javier, &javier)
	hola(&javier)
}
