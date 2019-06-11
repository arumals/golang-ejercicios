package main

import (
	"io/ioutil"
	"log"
)

func main() {
	b := []byte("Hola mundo!\n")
	err := ioutil.WriteFile("personal.txt", b, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
