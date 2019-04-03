package main

import (
	"fmt"
	"time"
)

func leerArchivo() string {
	time.Sleep(time.Second * 5)
	return "Datos del archivo"
}

func main() {
	go func() {
		datos := leerArchivo()
		fmt.Println(datos)
	}()
	fmt.Println("Continuar con la ejecución")
	time.Sleep(time.Second * 6)
}
