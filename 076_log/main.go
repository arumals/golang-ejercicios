package main

import (
	"errors"
	"log"
)

func main() {
	err := errors.New("Este es un error fatal de prueba")
	log.Fatal(err)
}
