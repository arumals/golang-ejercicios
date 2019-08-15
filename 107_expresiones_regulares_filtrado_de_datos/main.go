package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

func main() {

	// leer archivo html como bytes
	datosComoBytes, err := ioutil.ReadFile("lista.html")
	if err != nil {
		log.Fatal(err)
	}

	// preparar expresion regular de coincidencias
	coincidenciasRe := regexp.MustCompile(`<li>(.+)</li>`)

	// preparear expresion regular para limpiar html
	htmlRe := regexp.MustCompile(`<[^>]+>`)

	// extraer lineas donde aparezcan nombres
	nombres := coincidenciasRe.FindAllString(string(datosComoBytes), -1)

	// imprimir datos limpios
	for _, nombre := range nombres {
		nombreSinHtml := htmlRe.ReplaceAllString(nombre, "")
		fmt.Println(nombreSinHtml)
	}

}
