package main

import (
	"fmt"
	"time"
)

func main() {

	ahora := time.Now()
	fmt.Println("Fecha en este momento:")
	fmt.Println(ahora)
	fmt.Println("Dentro de un año:")
	fecha2 := ahora.Add(365 * 24 * time.Hour)
	fmt.Println(fecha2)
	fmt.Println("Hace un año:")
	fecha3 := ahora.Add(365 * 24 * time.Hour * -1)
	fmt.Println(fecha3)

}
