package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.TrimSpace("  remover espacio en ambos lados "))
	fmt.Println(strings.TrimLeft("¿Cuantos años tiene usted?", "¿"))
	fmt.Println(strings.TrimRight("¿Cuantos años tiene usted?", "?"))
}
