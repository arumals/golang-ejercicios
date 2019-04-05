package main

import (
	"fmt"
	"time"
)

func caballo1(can chan string) {
	time.Sleep(time.Second * 3)
	can <- "Caballo 1 completo la carrera"
}

func caballo2(can chan string) {
	time.Sleep(time.Second * 2)
	can <- "Caballo 2 completo la carrera"
}

func caballo3(can chan string) {
	time.Sleep(time.Second * 4)
	can <- "Caballo 3 completo la carrera"
}

func main() {

	can1 := make(chan string)
	can2 := make(chan string)
	can3 := make(chan string)

	go caballo1(can1)
	go caballo2(can2)
	go caballo3(can3)

	select {
	case ganador := <-can1:
		fmt.Println(ganador)
	case ganador := <-can2:
		fmt.Println(ganador)
	case ganador := <-can2:
		fmt.Println(ganador)
	case <-time.After(time.Second * 1):
		fmt.Println("Se excedio el tiempo de ejecución")
	}

}
