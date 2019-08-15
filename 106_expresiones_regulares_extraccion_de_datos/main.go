package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

func main() {

	// en donde vamos a capturar los nombres
	var nombres = make([]string, 3)

	// cargar contenido del html
	datosComoBytes, err := ioutil.ReadFile("lista.html")
	if err != nil {
		log.Fatal(err)
	}

	// preparar la expresion regular
	expReg := regexp.MustCompile(`(<span class="nombre">)([^<]+)(</span>)`)

	// ejecutar la busqueda de los indices
	todosLosIndices := expReg.FindAllSubmatchIndex(datosComoBytes, -1)

	// recorrer los resultados
	for _, loc := range todosLosIndices {
		//fmt.Println(loc)
		//fmt.Println(string(datosComoBytes[loc[0]:loc[1]]))
		//fmt.Println(string(datosComoBytes[loc[2]:loc[3]]))
		//fmt.Println(string(datosComoBytes[loc[4]:loc[5]]))
		//fmt.Println(string(datosComoBytes[loc[6]:loc[7]]))
		// capturar el nombre
		nombres = append(nombres, string(datosComoBytes[loc[4]:loc[5]]))
	}

	// imprimir los nombres
	fmt.Println(nombres)
}
