package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	texto, err := ioutil.ReadFile("archivo.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(texto)
}
