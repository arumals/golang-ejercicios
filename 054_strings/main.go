package main

import (
	"fmt"
	"strings"
)

func main() {
	cancion := "Nothing else matters"
	fmt.Println(strings.ToLower(cancion))
	fmt.Println(strings.ToUpper(cancion))
	fmt.Println(strings.Contains(cancion, "matters"))
}
