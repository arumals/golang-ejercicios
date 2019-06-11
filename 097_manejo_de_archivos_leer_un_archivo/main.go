package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// leer el arreglo de bytes del archivo
	datosComoBytes, err := ioutil.ReadFile("empleados.txt")
	if err != nil {
		log.Fatal(err)
	}
	// convertir el arreglo a string
	datosComoString := string(datosComoBytes)
	// imprimir el string
	fmt.Println(datosComoString)
}
