package main

import (
	"fmt"
	"strings"
)

func main() {
	serie := "Dr House especialista médico"
	if posicion := strings.Index(serie, "ista"); posicion != -1 {
		fmt.Println("La posicion de 'ista' es: ", posicion)
	}
	if posicion := strings.Index(serie, "mesta"); posicion == -1 {
		fmt.Println("'mesta' no se encuentra en el string")
	}
}
