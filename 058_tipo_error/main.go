package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("error generado durante la ejecución")
	if err != nil {
		fmt.Println(err)
	}
}
