package main

import "fmt"

func hola(s string) string {
	return "Hola " + s
}

func main() {
	fmt.Println(hola("Luis"))
}
