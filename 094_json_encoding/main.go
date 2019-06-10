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
	p := Pais{
		Nombre:     "Canada",
		Habitantes: 37314442,
		Capital:    "Ottawa",
		Idiomas:    []string{"Inglés", "Frances"},
	}
	datosJson, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(datosJson))
}
