package main

import (
	"fmt"
)

func main() {
	noticias := make(chan string, 3)
	noticias <- "Francia tendra elecciones el dia de hoy"
	noticias <- "En una votación cerrada Macron gana la presidencia"
	noticias <- "Macron es el nuevo presidente de Francia"
	close(noticias)
	for noticia := range noticias {
		fmt.Println(noticia)
	}
}
