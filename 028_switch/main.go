package main

import "fmt"

func main() {
	juguete := "cosa"
	switch juguete {
	case "persona":
		fmt.Println("El juguete es una persona")
	case "cosa":
		fmt.Println("El juguete es una cosa")
	case "animal":
		fmt.Println("El juguete es un animal")
	default:
		fmt.Println("El juguete es otra categoria")
	}
}
