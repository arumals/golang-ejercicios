package main

import (
	"fmt"
	"regexp"
)

func main() {
	// datos de entrada
	datosDeEntrada := []string{
		"2019-01-01 00:00:00a",
		"2019-01-01_ 00:00:00",
		"2019-01-01  00:00:00",
	}
	// preparar la expresion regular
	caracteresInvalidos := regexp.MustCompile("[^0-9 -:]")
	multiplesEspacios := regexp.MustCompile("[ ]{2,}")
	// sanitizar cada uno de los valores
	for _, dato := range datosDeEntrada {
		dato = caracteresInvalidos.ReplaceAllString(dato, "")
		dato = multiplesEspacios.ReplaceAllString(dato, " ")
		fmt.Println(dato)
	}

}
