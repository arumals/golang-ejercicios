package main

import (
	"fmt"
	"time"
)

func leerArchivo() string {
	time.Sleep(time.Second * 2)
	return "Datos del archivo"
}

func main() {
	miCanal := make(chan string)
	go func() {
		miCanal <- leerArchivo()
	}()
	fmt.Println(<-miCanal)
	fmt.Println("Continuar con la ejecución")
}
