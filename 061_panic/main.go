package main

import "fmt"

func main() {
	fmt.Println("Antes de iniciar pánico")
	panic("Iniciar pánico")
	fmt.Println("Esta línea de código no es alcanzada")
}
