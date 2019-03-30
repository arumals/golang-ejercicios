package main

import "fmt"

func main() {
	hola := "hola"
	fmt.Println(len(hola))
	fmt.Println(hola[0])
	fmt.Printf("%q\n", hola[0])
	fmt.Printf("%b\n", hola[0])
	nihao := "你好"
	fmt.Println(len(nihao))
	fmt.Printf("%q\n", nihao[0])
	fmt.Printf("%b\n", nihao[0])
}
