package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Ingresa tu nombre: ")
	nombre, _ := reader.ReadString('\n')
	nombre = strings.Replace(nombre, "\n", "", -1)
	fmt.Printf("Ingresa tu edad: ")
	edad, _ := reader.ReadString('\n')
	edad = strings.Replace(edad, "\n", "", -1)
	fmt.Println("-------------")
	fmt.Println("Nombre:", nombre)
	fmt.Println("Edad:", edad)
	fmt.Println("-------------")
}
