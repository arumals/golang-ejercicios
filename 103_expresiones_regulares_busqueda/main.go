package main

import (
	"fmt"
	"log"
	"regexp"
)

func main() {
	buscarTexto := "el perro"
	buscarEntTexto := "vuelve el perro arrepentido"
	existe, err := regexp.MatchString(buscarTexto, buscarEntTexto)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(existe)
}
