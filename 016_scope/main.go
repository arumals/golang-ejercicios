package main

import "fmt"

var var1 = "Este es el nivel 1"

func main() {
	var var2 = "Este es el nivel 2"
	{
		var var3 = "Este es el nivel 3"
	}
	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(var3)
}
