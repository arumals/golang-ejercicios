package main

import (
	"flag"
	"fmt"
)

func main() {
	nombre := flag.String("nombre", "", "El nombre de la persona")
	edad := flag.Int("edad", 18, "La edad de la persona")
	flag.Parse()
	fmt.Println("Tu nombre es:", *nombre)
	fmt.Println("Tu edad es:", *edad)
}
