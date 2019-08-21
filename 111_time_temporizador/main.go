package main

import (
	"fmt"
	"time"
)

func main() {

	// creamos el temporizador
	tdr := time.Tick(3 * time.Second)

	// recorremos el temporizador
	for horaActual := range tdr {
		fmt.Println("La hora es", horaActual)
	}

}
