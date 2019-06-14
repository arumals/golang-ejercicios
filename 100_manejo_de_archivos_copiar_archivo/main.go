package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// leer datos de origen
	origen, err := os.Open("origen.txt")
	if err != nil {
		log.Fatal(err)
	}
	// cierra el archivo "origen.txt" al terminar el programa
	defer origen.Close()
	// crea un nuevo archivo
	destino, err := os.OpenFile("destino.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	// cierra el archivo "destino.txt" al terminar programa
	defer destino.Close()
	// copiar datos
	resultadoCopia, err := io.Copy(destino, origen)
	if err != nil {
		log.Fatal(err)
	}
	//
	fmt.Println(resultadoCopia)
}
