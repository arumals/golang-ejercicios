package main

import "fmt"

func main() {
	razasDePerros := []string{"labrador", "poodle", "doberman", "shitzu", "beagle"}
	fmt.Println(razasDePerros)
	razasDePerros = append(razasDePerros[:2], razasDePerros[2+1:]...)
	fmt.Println(razasDePerros)
}
