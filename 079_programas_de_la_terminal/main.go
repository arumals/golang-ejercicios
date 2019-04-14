package main

import (
	"fmt"
	"os"
)

func main() {
	for k, v := range os.Args {
		fmt.Printf("Argumento %v: %v\n", k, v)
	}
}
