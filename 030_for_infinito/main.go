package main

import "fmt"

func desconectar() {
	fmt.Println("Se ha desconectado de la base de datos")
}

func leer() {
	fmt.Println("Se han leido los registros de la base de datos")
}

func actualizar() {
	fmt.Println("Se han actalizado registros de la base de datos")
}

func conectar() {
	fmt.Println("Se ha conectado a la base de datos")
}

func main() {
	conectar()
	defer desconectar()
	leer()
	actualizar()
}
