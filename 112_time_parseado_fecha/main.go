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

	fmt.Printf("El tipo es: %T \n", fecha)
	fmt.Printf("La fecha es: %v \n", fecha)

}
