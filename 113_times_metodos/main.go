package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	fechaComoString := "2015-05-15T10:12:11+06:00"

	fecha, err := time.Parse(time.RFC3339, fechaComoString)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("El año es:", fecha.Year())
	fmt.Println("El mes es:", fecha.Month())
	fmt.Println("El dia es:", fecha.Day())
	fmt.Println("La hora es:", fecha.Hour())
	fmt.Println("Los minutos son:", fecha.Minute())
	fmt.Println("Los segundos son:", fecha.Second())

}
