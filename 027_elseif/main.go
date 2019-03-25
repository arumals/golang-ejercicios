package main

import "fmt"

func main() {
	juguete := "cosa"
	if juguete == "persona" {
		fmt.Println("El objeto es una persona")
	} else if juguete == "cosa" {
		fmt.Println("El objeto es una cosa")
	} else if juguete == "animal" {
		fmt.Println("El objeto es un animal")
	} else {
		fmt.Println("El objeto es otra categoria")
	}
}
