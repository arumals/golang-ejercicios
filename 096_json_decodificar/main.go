package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Pais struct {
	Nombre     string
	Habitantes int
	Capital    string
	Idiomas    []string
}

func main() {
	datosJson := `{"nombre":"Canada","habitantes":37314442,"capital":"Ottawa","idiomas":["Inglés","Frances"]}`
	p := Pais{}
	err := json.Unmarshal([]byte(datosJson), &p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", p)
}
