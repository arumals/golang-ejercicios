package main

import "fmt"

func main() {
	diasDeLaSemanaEnIngles := map[string]string{
		"lunes":     "Monday",
		"martes":    "Tuesday",
		"miercoles": "Wednesday",
		"jueves":    "Thursday",
		"sabado":    "Saturday",
		"domingo":   "Sunday",
	}
	fmt.Println(diasDeLaSemanaEnIngles)
	delete(diasDeLaSemanaEnIngles, "domingo")
	fmt.Println(diasDeLaSemanaEnIngles)
}
