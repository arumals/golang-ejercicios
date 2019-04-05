package main

import "fmt"

func soloLectura(noticias <-chan string) {
	fmt.Println(<-noticias)
}

func soloEscritura(noticias chan<- string) {
	noticias <- "Noticia 2"
}

func lecturaEscritura(noticias chan string) {
	noticias <- "Noticia 3"
	fmt.Println(<-noticias)
}

func main() {
	nts := make(chan string, 2)
	nts <- "Noticia 1"
	soloLectura(nts)
	soloEscritura(nts)
	fmt.Println(<-nts)
	lecturaEscritura(nts)
}
