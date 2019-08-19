package main

import (
	"fmt"
	"time"
)

func leerApi() string {
	time.Sleep(5 * time.Second)
	return "respuesta del api"
}

func main() {

	c1 := make(chan string, 1)

	go func() {
		c1 <- leerApi()
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout")
	}
}
