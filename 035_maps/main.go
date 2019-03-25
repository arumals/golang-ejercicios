package main

import "fmt"

func main() {
	var diasDeLaSemana = make(map[string]int)
	diasDeLaSemana["lunes"] = 1
	diasDeLaSemana["martes"] = 2
	diasDeLaSemana["miercoles"] = 3
	diasDeLaSemana["jueves"] = 4
	diasDeLaSemana["viernes"] = 5
	diasDeLaSemana["sabado"] = 6
	diasDeLaSemana["domingo"] = 7
	fmt.Println(diasDeLaSemana)
}
