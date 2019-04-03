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
	datos := leerArchivo()
	fmt.Println(datos)
}
