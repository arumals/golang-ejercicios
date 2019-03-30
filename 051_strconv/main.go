package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println("Contando hasta 100 : " + strconv.Itoa(i))
	}
}
