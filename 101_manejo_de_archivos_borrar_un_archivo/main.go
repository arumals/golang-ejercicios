package main

import (
	"log"
	"os"
)

func main() {
	// borrar el archivo archivoBorrable.txt
	err := os.Remove("archivoBorrable.txt")
	if err != nil {
		log.Fatal(err)
	}
}
