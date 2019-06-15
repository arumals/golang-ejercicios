package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Credenciales struct {
	Usuario  string `json:"usuario"`
	Password string `json:"password"`
}

func main() {
	// abrir archivo "configuracion.json"
	manejadorDeArchivo, err := ioutil.ReadFile("credenciales.json")
	if err != nil {
		log.Fatal(err)
	}
	// preparar contenedor de las credenciales
	c := Credenciales{}
	// decodificar el contenido del json sobre la estructura
	err = json.Unmarshal(manejadorDeArchivo, &c)
	if err != nil {
		log.Fatal(err)
	}
	// validar credenciales
	if c.Usuario == "luis" && c.Password == "passwordDeLuis" {
		fmt.Println("Las credenciales son correctas")
	} else {
		fmt.Println("Las credenciales son incorrectas")
	}
}
