package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	archivos, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, archivo := range archivos {
		fmt.Println("Nombre:", archivo.Name())
		fmt.Println("Tamaño:", archivo.Size())
		fmt.Println("Modo:", archivo.Mode())
		fmt.Println("Ultima modificación:", archivo.ModTime())
		fmt.Println("Es directorio?:", archivo.IsDir())
		fmt.Println("-----------------------------------------")
	}
}
