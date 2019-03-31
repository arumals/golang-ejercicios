package main

import "fmt"

func main() {
	edadMinima := "18"
	err := fmt.Errorf("La edad debe ser mayor de de %v años", edadMinima)
	if err != nil {
		fmt.Println(err)
	}
}
