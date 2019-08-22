package main

import (
	"fmt"
	"time"
)

func main() {

	ahora := time.Now()
	proxSemana := ahora.Add(time.Hour * 24 * 7)

	fmt.Println(ahora.Equal(proxSemana))
	fmt.Println(ahora.Before(proxSemana))
	fmt.Println(ahora.After(proxSemana))

}
