package main

import (
	"fmt"
	"time"
)

func leerArchivo1() string {
	time.Sleep(time.Second * 3)
	return "Datos del archivo 1"
}

func leerArchivo2() string {
	time.Sleep(time.Second * 3)
	return "Datos del archivo 2"
}

func main() {
	miCanal := make(chan string, 2)
	go func() {
		miCanal <- leerArchivo1()
	}()
	go func() {
		miCanal <- leerArchivo2()
	}()
	fmt.Println(<-miCanal)
	fmt.Println(<-miCanal)
}
